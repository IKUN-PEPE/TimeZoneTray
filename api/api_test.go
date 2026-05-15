package api

import (
	"strings"
	"testing"
)

func TestGetLatLon(t *testing.T) {
	lat, lon, err := GetLatLon("10001")
	if err != nil {
		t.Fatal(err)
	}
	expectedLat := "40.7484"
	expectedLon := "-73.9967"
	if lat != expectedLat || lon != expectedLon {
		t.Errorf("expected %s, %s; got %s, %s", expectedLat, expectedLon, lat, lon)
	}
}

func TestGetTimeByZip(t *testing.T) {
	timeStr, err := GetTimeByZip("10001")
	if err != nil {
		t.Fatal(err)
	}
	if timeStr == "" {
		t.Errorf("expected non-empty time string")
	}
	// Time zone for 10001 is typically "America/New_York", so check if it contains that.
	if !strings.Contains(timeStr, "America/New_York") {
		t.Errorf("expected time string to contain time zone 'America/New_York', got: %s", timeStr)
	}
}
