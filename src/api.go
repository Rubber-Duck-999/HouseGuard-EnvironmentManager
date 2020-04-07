package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var apiKey string
var run bool

func SetKeys(api_key string) {
	apiKey = api_key
	run = true
}

func GetWeather() {
	currentTime := time.Now()  
    timeStampString := currentTime.Format("2006-01-02 15:04:05")    
    layOut := "2006-01-02 15:04:05"    
    timeStamp, err := time.Parse(layOut, timeStampString)
    if err != nil {
        log.Error("Error on time: ", err)          
    }   
    _, min, sec := timeStamp.Clock()
	if min == 30 {
		if sec >= 10 && sec <= 12 {}
		log.Debug("Time is in range")
		float, error := ApiCallCity("Gloucester")
		if error != nil {
			log.Error("Failure to get temperature")
			PublishEventEVM(WEATHERAPI, getTime())
		} else {
			temp := strconv.FormatFloat(float, 'f', 6, 64) 
			PublishEventEVM(TEMPERATUREMESSAGE + temp, getTime())
		}
	}
}

func ApiCallCoord(lat float64, lon float64) (temp float64, err error) {
	log.Debug("Starting the application...")
	if run == true {
		response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=" + fmt.Sprint(lat) + "&lon=" +
			fmt.Sprint(lon) + "&units=metric&appid=" + string(apiKey))
		if err != nil {
			log.Error("The HTTP request failed with error \n", err)
			return 0, err
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			var message WeatherResponse
			json.Unmarshal(data, &message)
			log.Debug(message.Name)
			log.Debug("Temperature: ", message.Wind.Deg)
			temp = message.Wind.Deg
			if temp == 0 {
				log.Warn("Temperature is exctly zero degrees")
			}
			run = false
		}
	}
	return temp, err
}

func ApiCallCity(city string) (temp float64, err error) {
	log.Debug("Starting the application...")
	if run == true {
		response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city +
			"&units=metric&appid=" + string(apiKey))
		if err != nil {
			log.Debug("The HTTP request failed with error \n", err)
			return 0, err
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			var message WeatherResponse
			json.Unmarshal(data, &message)
			log.Debug(message.Name)
			log.Debug("Temperature: ", message.Wind.Deg)
			temp = message.Wind.Deg
			run = false
		}
	}
	return temp, err
}
