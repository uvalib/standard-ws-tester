package tests

import (
	"net/http"
	"testing"
)

//
// healthcheck tests
//

func TestHealthCheck(t *testing.T) {
	expected := http.StatusOK
	status, _ := HealthCheck(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	// enumerate each section and ensure that the health is reported as true
	//for k, v := range response {
	//	if v.Healthy != true {
	//		t.Fatalf("Expected healthy for %s\n", k)
	//	}
	//}
}

//
// end of file
//
