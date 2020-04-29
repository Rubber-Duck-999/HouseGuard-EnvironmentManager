// rabbitmq_test.go

package main

import (
	"time"
	"testing"
)

func TestCheckKeys(t *testing.T) {
	SetKeys("cheese")
	if apiKey != "cheese" {
		t.Error("Failure")
	}
}

func TestCheckInCorrectKeys(t *testing.T) {
	SetKeys("cheese")
	file := "./EVM.yml"
	var data ConfigTypes
	if Exists(file) {
		GetData(&data, file)
	}
	SetPassword(data.Settings.Pass)
	err := SetConnection()
	if err == nil {
		t.Error("Failure")
	}

	now := time.Now()
	m := now.Minute()
	weather_minute = m
	if apiKey != "cheese" {
		t.Error("Failure")
	}
	GetWeather()
	if float > 0.00 {
		t.Error("Failure")
	}
}

func TestCheckCorrectKeys(t *testing.T) {
	file := "../EVM.yml"
	var data ConfigTypes
	if Exists(file) {
		GetData(&data, file)
	}
	SetPassword(data.Settings.Pass)
	SetKeys(data.Settings.Key)
	err := SetConnection()
	if err != nil {
		t.Error("Failure")
	}

	now := time.Now()
	m := now.Minute()
	weather_minute = m
	if apiKey == "cheese" {
		t.Error("Failure")
	}
	GetWeather()
	if float < 0.01 {
		t.Error("Failure")
	}
}
