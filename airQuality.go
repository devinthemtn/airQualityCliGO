package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type weatherAPI struct {
	DateObserved  string
	HourObserved  int
	LocalTimeZone string
	ReportingArea string
	StateCode     string
	Latitude      float32
	Longitude     float32
	ParameterName string
	AQI           int
	Category      weatherAPICategory
}

type weatherAPICategory struct {
	Number int
	Name   string
}

func main() {
	args := os.Args[1:]

	var zip string
	if len(args) > 0 {
		zip = args[0]
	} else {
		zip = "59865"
	}

	resp, err := http.Get("http://www.airnowapi.org/aq/observation/zipCode/current/?format=application/json&zipCode=" + zip + "&distance=90&API_KEY=2C4804CE-3667-4B35-8D78-A175206A29DA")
	if err != nil {
		fmt.Println("error", err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	var weatherAPIObj []weatherAPI
	json.Unmarshal(bodyBytes, &weatherAPIObj)

	// fmt.Println(weatherAPIObj[0])
	fmt.Printf("Location: %s, %s \n", weatherAPIObj[0].ReportingArea, weatherAPIObj[0].StateCode)
	fmt.Printf("Air quality: %d (%s) \n", weatherAPIObj[0].AQI, weatherAPIObj[0].Category.Name)
}
