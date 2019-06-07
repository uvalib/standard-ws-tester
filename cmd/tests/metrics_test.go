package tests

import (
	"net/http"
	"strings"
	"testing"
)

//
// metrics tests
//

func TestMetricsCheck(t *testing.T) {
	expected := http.StatusOK
	status, metrics := MetricsCheck(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if emptyField(metrics) {
		t.Fatalf("Expected non-empty metrics info\n")
	}

	if strings.Contains(metrics, "go_goroutines") == false {
		t.Fatalf("Expected go_goroutines value in metrics info\n")
	}
}

//
// end of file
//