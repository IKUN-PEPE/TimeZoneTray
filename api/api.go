package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ZippoResponse represents the JSON structure returned by api.zippopotam.us
type ZippoResponse struct {
	PostCode      string `json:"post code"`
	Country       string `json:"country"`
	CountryAbbrev string `json:"country abbreviation"`
	Places        []struct {
		PlaceName string `json:"place name"`
		Longitude string `json:"longitude"`
		State     string `json:"state"`
		StateAbr  string `json:"state abbreviation"`
		Latitude  string `json:"latitude"`
	} `json:"places"`
}

// GetLatLon fetches latitude and longitude for a given US zip code using Zippopotam API.
func GetLatLon(zip string) (string, string, error) {
	url := fmt.Sprintf("http://api.zippopotam.us/us/%s", zip)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("api returned status: %s", resp.Status)
	}

	var zippo ZippoResponse
	if err := json.NewDecoder(resp.Body).Decode(&zippo); err != nil {
		return "", "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(zippo.Places) == 0 {
		return "", "", fmt.Errorf("no places found for zip code %s", zip)
	}

	// Return the latitude and longitude of the first place
	return zippo.Places[0].Latitude, zippo.Places[0].Longitude, nil
}
