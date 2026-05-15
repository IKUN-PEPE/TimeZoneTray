package api

import "testing"

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
