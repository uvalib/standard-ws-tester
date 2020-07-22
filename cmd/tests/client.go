package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var debugHTTP = false
var serviceTimeout = 15
var client = newHttpClient(serviceTimeout)

//
// newHttpClient -- create a new client
//
func newHttpClient(timeout int) *http.Client {

	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}

//
// HealthCheck -- calls the service health check method
//
func HealthCheck(endpoint string) (int, map[string]HealthCheckResult) {

	url := fmt.Sprintf("%s/healthcheck", endpoint)

	status, body := httpGet(url)
	if status != http.StatusOK {
		return status, nil
	}

	// cos we dont really know the field names ahead of time
	r := make(map[string]HealthCheckResult)

	err := json.Unmarshal(body, &r)
	if err != nil {
		fmt.Printf("ERROR: unmarshal of [%s] returns: %s\n", body, err)
		return http.StatusInternalServerError, nil
	}

	return status, r
}

//
// VersionCheck -- calls the service version check method
//
func VersionCheck(endpoint string) (int, string) {

	url := fmt.Sprintf("%s/version", endpoint)

	status, body := httpGet(url)
	if status != http.StatusOK {
		return status, ""
	}

	r := VersionResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		fmt.Printf("ERROR: unmarshal of [%s] returns: %s\n", body, err)
		return http.StatusInternalServerError, ""
	}

	return status, r.Build
}

//
// MetricsCheck -- calls the service metrics method
//
func MetricsCheck(endpoint string) (int, string) {

	url := fmt.Sprintf("%s/metrics", endpoint)

	status, body := httpGet(url)
	if status != http.StatusOK {
		return status, ""
	}

	return status, string(body)
}

func httpGet(url string) (int, []byte) {

	//fmt.Printf( "%s\n", url )

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("ERROR: request (%s) returns: %s\n", url, err)
		return http.StatusInternalServerError, nil
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "")

	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("ERROR: request (%s) returns: %s\n", url, err)
		return http.StatusInternalServerError, nil
	}

	defer io.Copy(ioutil.Discard, response.Body)
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: request (%s) returns: HTTP %d\n", url, response.StatusCode)
		return response.StatusCode, nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("ERROR: request (%s) returns: %s\n", url, err)
		return http.StatusInternalServerError, nil
	}

	return response.StatusCode, body
}

//
// end of file
//
