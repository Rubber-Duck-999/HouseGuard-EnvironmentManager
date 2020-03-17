package main

type coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type weather []struct {
	Id int             `json:"id"`
	Main string        `json:"main"`
	Description string `json:"description"`
	Icon string        `json:"icon"`
}

type mainWeather struct {
	Temp     int `json:"temp"`
	Like     int `json:"feels_like"`
	Min      int `json:"temp_min"`
	Max      int `json:"temp_max"`
	Pressure int `json:"pressure"`
	Humidity int `json:"humidity"`
}

type wind struct {
	Speed int `json:"speed"`
	Deg int   `json:"deg"`
}

type cloud struct {
	All int `json:"all"`
}

type system struct {
	Type    int `json:"type"`
	Id      int `json:"id"`
	Message float32 `json:"message"`
	Country string  `json:"country"`
	Sunrise int64   `json:"sunrise"`
	Sunset  int64   `json:"sunset"`
}

type WeatherResponse struct {
	Coord coordinates `json:"coord"`
	Weather weather   `json:"weather"`
	Base string       `json:"base"`
	Main mainWeather  `json:"main"`
	Visibility int    `json:"visibility"`
	Wind wind         `json:"wind"`
	Clouds cloud      `json:"clouds"`
	DT int64          `json:"dt"`
	Sys system        `json:"sys"`
	Timezone int64    `json:"timezone"`
	ID int64          `json:"id"`
	Name string       `json:"name"`
	Cod int           `json:"cod"`
}

type FailureMessage struct {
	Time         string `json:"time"`
	Failure_type string `json:"type"`
	Severity     int    `json:"severity"`
}

type MotionDetected struct {
	Time string
}

type ConfigTypes struct {
	Settings struct {
		Key     string `yaml:"key"`
	} `yaml:"settings"`
}

type EventEVM struct {
	Component string
	Message   string
	Time      string
	Severity  int
}

type MapMessage struct {
	message     string
	routing_key string
	time        string
	valid       bool
}

const WEATHER string = "Weather"
const FAILURECOMPONENT string = "Failure.Component"
const MOTIONDETECTED string = "Motion.Detected"
const MOTIONRESPONSE string = "Motion.Response"
const ISSUENOTICE string = "Issue.Notice"
const EVENTEVM string = "Event.EVM"
const EXCHANGENAME string = "topics"
const EXCHANGETYPE string = "topic"
const TIMEFORMAT string = "20060102150405"
const CAMERAMONITOR string = "CM"
const COMPONENT string = "EVM"
const UPDATESTATEERROR string = "We have received a brand new state update"
const SERVERERROR string = "Server is failing to send"
const STATEUPDATESEVERITY int = 2
const SERVERSEVERITY int = 4
const FAILURECONVERT string = "Failed to convert"
const FAILUREPUBLISH string = "Failed to publish"

var SubscribedMessagesMap map[uint32]*MapMessage
var key_id uint32 = 0
