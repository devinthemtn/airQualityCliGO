package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.airnowapi.org/aq/observation/zipCode/current/?format=application/json&zipCode=59865&distance=50&API_KEY=2C4804CE-3667-4B35-8D78-A175206A29DA")
	if err != nil {
		fmt.Println("error", err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println("API response: " + bodyString)
}
