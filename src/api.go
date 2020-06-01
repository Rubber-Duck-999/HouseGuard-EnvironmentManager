package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var apiKey string
var weather_minute int

func init() {
	apiKey = "N/A"
	weather_minute = 60
}

func SetKeys(api_key string) {
	apiKey = api_key
}

func GetWeather() {
	temporary, error := ApiCallCity("Gloucester")
	if conn != nil {
		if error != nil {
			log.Error("Failure to get temperature")
			PublishEventEVM(WEATHERAPI, getTime(), "EVM2")
		} else {
			current_temp = strconv.FormatFloat(temporary, 'f', 6, 64) 
			PublishEventEVM(TEMPERATUREMESSAGE + current_temp, getTime(), "EVM3")
			_statusEVM.CurrentTemperature = temporary
		}
	}
}

func ApiCallCity(city string) (temp float64, err error) {
	log.Debug("Starting the application...")
	err = nil
	if apiKey != "N/A" {
		response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city +
		"&units=metric&appid=" + string(apiKey))
		if err != nil {
			log.Error("The HTTP request failed with error \n", err)
			return 0, err
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			var message WeatherResponse
			json.Unmarshal(data, &message)
			log.Debug(message.Name)
			temp = message.Wind.Deg / 10
		}
	}
	return temp, err
}
