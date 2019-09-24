package tests

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var debugHTTP = false
var serviceTimeout = 5

//
// HealthCheck -- calls the service health check method
//
func HealthCheck(endpoint string) (int, map[string] HealthCheckResult ) {

	url := fmt.Sprintf("%s/healthcheck", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Set("Accept", "application/json").
		Set("Content-Type", "").
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, nil
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	// cos we dont really know the field names ahead of time
	m := make( map[string] HealthCheckResult )

	err := json.Unmarshal([]byte(body), &m)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return resp.StatusCode, m
}

//
// VersionCheck -- calls the service version check method
//
func VersionCheck(endpoint string) (int, string) {

	url := fmt.Sprintf("%s/version", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Set("Accept", "application/json").
		Set("Content-Type", "").
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, ""
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := VersionResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, ""
	}

	return resp.StatusCode, r.Build
}

//
// MetricsCheck -- calls the service metrics method
//
func MetricsCheck(endpoint string) (int, string) {

	url := fmt.Sprintf("%s/metrics", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Set("Accept", "*/*").
		Set("Content-Type", "").
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, ""
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	return resp.StatusCode, body
}

//
// end of file
//
