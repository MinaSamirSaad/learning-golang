package api_test

import (
	"testing"

	"test.com/go/cryptomasters/api"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("BTC")
	if err != nil {
		t.Errorf("Error in getting rate %v", err)
	}
}
