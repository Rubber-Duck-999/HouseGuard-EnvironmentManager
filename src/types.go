package main

type coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type weather []struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
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
	Speed int     `json:"speed"`
	Deg   float64 `json:"deg"`
}

type cloud struct {
	All int `json:"all"`
}

type system struct {
	Type    int     `json:"type"`
	Id      int     `json:"id"`
	Message float32 `json:"message"`
	Country string  `json:"country"`
	Sunrise int64   `json:"sunrise"`
	Sunset  int64   `json:"sunset"`
}

type WeatherResponse struct {
	Coord      coordinates `json:"coord"`
	Weather    weather     `json:"weather"`
	Base       string      `json:"base"`
	Main       mainWeather `json:"main"`
	Visibility int         `json:"visibility"`
	Wind       wind        `json:"wind"`
	Clouds     cloud       `json:"clouds"`
	DT         int64       `json:"dt"`
	Sys        system      `json:"sys"`
	Timezone   int64       `json:"timezone"`
	ID         int64       `json:"id"`
	Name       string      `json:"name"`
	Cod        int         `json:"cod"`
}

type FailureMessage struct {
	Time     string `json:"time"`
	Severity int    `json:"severity"`
}

type MotionResponse struct {
	File string `json:"file"`
	Time string `json:"time"`
	Severity int `json:"severity"`
}

type MotionDetected struct {
	File string
	Time string
}

type ConfigTypes struct {
	Settings struct {
		Key string `yaml:"Key"`
		Pass string `yaml:"Pass"`
		Sheet string `yaml:"Sheet"`
	} `yaml:"settings"`
}

type EventEVM struct {
	Component string
	Message   string
	Time      string
}

type MapMessage struct {
	message     string
	routing_key string
	time        string
	valid       bool
}

// Status Messages
type StatusDBM struct {
	DailyEvents int
	TotalEvents int
	CommonEvent string
	DailyDataRequests int
}

type StatusSYP struct {
	HighestUsage int
	MemoryLeft int
}

type StatusFH struct {
	DailyFaults int
	CommonFaults string
}

type StatusNAC struct {
	DevicesActive int
	DailyBlockedDevices int
	DailyUnknownDevices int
	DailyAllowedDevices int
	TimeEscConnected string
}

type StatusEVM struct {
	DailyImagesTaken int
	CurrentTemperature int
	LastMotionDetected string
}

type StatusUP struct {
	LastAccessGranted string
	LastAccessBlocked string
	CurrentAlarmState string
	LastUser string
}

const STATUSDBM string = "Status.DBM"
const STATUSSYP string = "Status.SYP"
const STATUSFH  string = "Status.FH"
const STATUSNAC string = "Status.NAC"
const STATUSUP  string = "Status.UP" 
const STATUSALL string = "Status.*"
const STATUSREQUESTUP  string = "Status.Request.UP"
const STATUSREQUESTDBM string = "Status.Request.DBM"
//
const FAILURECOMPONENT string = "Failure.Component"
const MOTIONDETECTED string = "Motion.Detected"
const MOTIONRESPONSE string = "Motion.Response"
const EVENTEVM string = "Event.EVM"
//
const EXCHANGENAME string = "topics"
const EXCHANGETYPE string = "topic"
const TIMEFORMAT string = "2006/01/02 15:04:05"
const CAMERAMONITOR string = "CM"
const COMPONENT string = "EVM"
const SERVERSEVERITY int = 6
const EVENTTEMP int = 1
const MAXMESSAGES int = 10
const SENSORNETWORKDOWN string = "Lost connection to sensor"
const WEATHERAPI string = "Weather api not accessible"
const TEMPERATUREMESSAGE string = "Temperature ="
const FAILURECONVERT string = "Failed to convert"
const FAILUREPUBLISH string = "Failed to publish"

var SubscribedMessagesMap map[uint32]*MapMessage
var key_id uint32 = 0
var device_id uint32 = 0
