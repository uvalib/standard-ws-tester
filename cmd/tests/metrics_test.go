package tests

import (
	"net/http"
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

	if emptyField(metrics) == true {
		t.Fatalf("Expected non-empty metrics info\n")
	}

	//if strings.Contains(metrics, "# TYPE") == false && strings.Contains(metrics, "# HELP") == false {
	//	t.Fatalf("Expected \"# TYPE\" or \"# HELP\" value in metrics info\n")
	//}
}

//
// end of file
//
