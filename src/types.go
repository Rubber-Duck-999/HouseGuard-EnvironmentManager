package main

type FailureMessage struct {
	Time     string `json:"time"`
	Severity int    `json:"severity"`
}

type MotionResponse struct {
	File     string `json:"file"`
	Time     string `json:"time"`
	Severity int    `json:"severity"`
}

type MotionDetected struct {
	File string
	Time string
}

type ConfigTypes struct {
	Settings struct {
		Key      string `yaml:"Key"`
		Pass     string `yaml:"Pass"`
		Sheet    string `yaml:"Sheet"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
		ClientID string `yaml:"ClientId"`
		Endpoint string `yaml:"Endpoint"`
	} `yaml:"settings"`
}

type EventEVM struct {
	Component   string `json:"component"`
	Time        string `json:"time"`
	EventTypeId string `json:"event_type_id"`
}

type MapMessage struct {
	message     string
	routing_key string
	time        string
	valid       bool
}

// Alarm Event
type AlarmEvent struct {
	User   string `json:"user"`
	State  string `json:"state"`
}

// Daily Status
type DailyStatus struct {
	CreatedDate string
	Allowed     int
	Blocked     int
	Unknown     int
	TotalEvents int
	CommonEvent string
	TotalFaults int
	CommonFault string
}

// Status
type Status struct {
	CreatedDate    string
	MotionDetected string
	AccessGranted  string
	AccessDenied   string
	LastFault      string
	LastUser       string
	CPUTemp        int
	CPUUsage       int
	Memory         int
}

// Status Messages
type StatusDBM struct {
	DailyEvents       int    `json:"_dailyEvents"`
	TotalEvents       int    `json:"_totalEvents"`
	CommonEvent       string `json:"_commonEvent"`
	DailyDataRequests int    `json:"_dailyDataRequests"`
}

type StatusSYP struct {
	Temperature  int `json:"temperature"`
	MemoryLeft   int `json:"memory_left"`
	HighestUsage int `json:"highest_usage"`
}

type StatusFH struct {
	DailyFaults  int    `json:"daily_faults"`
	CommonFaults string `json:"common_faults"`
}

type StatusNAC struct {
	DevicesActive       int    `json:"devices_active"`
	DailyBlockedDevices int    `json:"blocked"`
	DailyUnknownDevices int    `json:"unknown"`
	DailyAllowedDevices int    `json:"allowed"`
	TimeEscConnected    string `json:"time"`
}

type StatusEVM struct {
	DailyImagesTaken   int
	LastMotionDetected string
}

type StatusUP struct {
	LastAccessGranted string `json:"_accessGranted"`
	LastAccessBlocked string `json:"_accessblocked"`
	CurrentAlarmState string `json:"_state"`
	LastUser          string `json:"_user"`
}

const STATUSDBM string = "Status.DBM"
const STATUSSYP string = "Status.SYP"
const STATUSFH string = "Status.FH"
const STATUSNAC string = "Status.NAC"
const STATUSUP string = "Status.UP"
const STATUSALL string = "Status.*"
const STATUSREQUESTUP string = "Status.Request.UP"
const STATUSREQUESTDBM string = "Status.Request.DBM"

//
const FAILURECOMPONENT string = "Failure.Component"
const MOTIONDETECTED string = "Motion.Detected"
const MOTIONRESPONSE string = "Motion.Response"
const EVENTEVM string = "Event.EVM"
const ALARMEVENT string = "Alarm.Event"
//
const EXCHANGENAME string = "topics"
const EXCHANGETYPE string = "topic"
const TIMEFORMAT string = "2006/01/02 15:04:05"
const CAMERAMONITOR string = "CM"
const COMPONENT string = "EVM"
const SERVERSEVERITY int = 6
const MAXMESSAGES int = 100
const FAILURECONVERT string = "Failed to convert"
const FAILUREPUBLISH string = "Failed to publish"

var SubscribedMessagesMap map[uint32]*MapMessage
var key_id uint32 = 0
var device_id uint32 = 0
