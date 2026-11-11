package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type weatherAPI struct {
	DateObserved       string `json:"dateObserved"`
	HourObserved       string `json:"hourObserved"`
	LocalTimeZone      string `json:"localTimeZone"`
	ReportingAreaName  string `json:"reportingAreaName"`
	SiteID             string `json:"siteID"`
	SiteName           string `json:"siteName"`
	ParameterName      string `json:"parameterName"`
	NowcastAQI         int    `json:"nowcastAQI"`
	AQICategoryName    string `json:"aqiCategoryName"`
	ReportingAgency    string `json:"reportingAgency"`
	LookupBehavior     string `json:"lookupBehavior"`
	ConsideredMonitors string `json:"consideredMonitors"`
	LookupBoundary     string `json:"lookupBoundary"`
}

type weatherAPICategory struct {
	Number int
	Name   string
}

// loadDotEnv reads KEY=VALUE lines from path into the environment.
// Variables already set in the environment take precedence.
func loadDotEnv(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(strings.TrimPrefix(key, "export "))
		value = strings.Trim(strings.TrimSpace(value), `"'`)
		if _, exists := os.LookupEnv(key); !exists {
			os.Setenv(key, value)
		}
	}
}

func main() {
	args := os.Args[1:]

	var zip string
	if len(args) > 0 {
		zip = args[0]
	} else {
		zip = "59865"
	}

	loadDotEnv(".env")
	apiKey := os.Getenv("AIRNOW_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "AIRNOW_API_KEY is not set (add it to .env or export it)")
		os.Exit(1)
	}

	q := url.Values{}
	q.Set("format", "application/json")
	q.Set("zipcode", zip)
	q.Set("distance", "90")
	q.Set("API_KEY", apiKey)

	resp, err := http.Get("https://www.airnowapi.org/aq/observation/current/ziplatlong/?" + q.Encode())
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	var weatherAPIObj []weatherAPI
	json.Unmarshal(bodyBytes, &weatherAPIObj)

	for _, entry := range weatherAPIObj {
		fmt.Printf("Location: %s, %s, %s \n", entry.ReportingAreaName, entry.DateObserved, entry.HourObserved)
		fmt.Printf("Air quality (%s): %d (%s) \n", entry.ParameterName, entry.NowcastAQI, entry.ParameterName)
	}
}
