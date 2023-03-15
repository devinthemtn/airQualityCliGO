package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.airnowapi.org/aq/observation/zipCode/current/?format=application/json&zipCode=59865&distance=50&API_KEY=REDACTED")
	if err != nil {
		fmt.Println("error", err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println("API response: " + bodyString)
}
