package main

import (
	"encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	
	log "github.com/sirupsen/logrus"
)

var apiKey string

func SetKeys(api_key string) {
	apiKey = api_key
}

func ApiCall(lat float64, lon float64) {
	log.Debug("Starting the application...")
	run := false
	if run == true {
		response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=" + fmt.Sprint(lat) + "&lon=" +
								fmt.Sprint(lon) + "&units=metric&appid=" + apiKey)
		if err != nil {
			log.Debug("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			var message WeatherResponse
			json.Unmarshal(data, &message)
			log.Debug(message.Name)
		}
	}
}