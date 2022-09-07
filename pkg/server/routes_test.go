package server

import "testing"

func TestGetEndpoints(t *testing.T) {
	endpointsMap := GetEndpoints()
	expected := []string{"/hello", "/"}

	for _, e := range expected {
		_, exists := endpointsMap[e]
		if !exists {
			t.Errorf("%s not found among endpoints", e)
		}
	}
}
