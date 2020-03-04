package rabbitmq

type FailureMessage struct {
	Time         string `json:"time"`
	Failure_type string `json:"type"`
	Severity     int    `json:"severity"`
}

type MotionDetected struct {
	Time string
}

type EventEVM struct {
	Component    string
	Message      string
	Time         string
	Severity     int
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
