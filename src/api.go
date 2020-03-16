package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func ApiCall() {
	fmt.Println("Starting the application...")
	run := false
	if run == true {
		response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=51.86&lon=2.23&units=metric&appid=")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	}
}