package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

// TimeAPIResponse represents the JSON structure returned by timeapi.io
type TimeAPIResponse struct {
	CurrentLocalTime string `json:"currentLocalTime"`
	TimeZone         string `json:"timeZone"`
}

// GetLatLon fetches latitude and longitude for a given US zip code using Zippopotam API.
func GetLatLon(ctx context.Context, zip string) (string, string, error) {
	url := fmt.Sprintf("http://api.zippopotam.us/us/%s", zip)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", "", fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
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

// GetTimeByCoord fetches the current time for a given latitude and longitude using TimeAPI.
func GetTimeByCoord(ctx context.Context, lat, lon string) (string, error) {
	url := fmt.Sprintf("https://timeapi.io/api/TimeZone/coordinate?latitude=%s&longitude=%s", lat, lon)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch time: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("time api returned status: %s", resp.Status)
	}

	var timeResp TimeAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&timeResp); err != nil {
		return "", fmt.Errorf("failed to decode time response: %w", err)
	}

	// The time format from TimeAPI often includes fractional seconds (e.g., 2026-05-14T23:23:15.6852017)
	// We only care about the part before the dot for "2006-01-02T15:04:05" format
	dateTimeParts := strings.Split(timeResp.CurrentLocalTime, ".")
	layout := "2006-01-02T15:04:05"

	t, err := time.Parse(layout, dateTimeParts[0])
	if err != nil {
		return "", fmt.Errorf("failed to parse date time '%s': %w", timeResp.CurrentLocalTime, err)
	}

	// Format to "YYYY-MM-DD HH:mm:ss (TimeZoneName)"
	formattedTime := fmt.Sprintf("%s (%s)", t.Format("2006-01-02 15:04:05"), timeResp.TimeZone)

	return formattedTime, nil
}

// GetTimeByZip fetches current time for a given zip code by first getting coordinates then calling TimeAPI.
func GetTimeByZip(ctx context.Context, zip string) (string, error) {
	lat, lon, err := GetLatLon(ctx, zip)
	if err != nil {
		return "", err
	}
	return GetTimeByCoord(ctx, lat, lon)
}
