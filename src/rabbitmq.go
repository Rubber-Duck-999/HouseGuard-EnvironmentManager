package main

import (
	"time"

	"github.com/clarketm/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel
var password string
var init_err error
var tempSet bool
var current_temp string
var float float64
var _year int 
var _month time.Month
var _day int
var _messages_sent int
//Status
var _statusDBM StatusDBM
var _statusSYP StatusSYP
var _statusFH  StatusFH 
var _statusNAC StatusNAC
var _statusEVM StatusEVM
var _statusUP  StatusUP
//


func init() {
	log.Trace("Initialised rabbitmq package")
	current_temp = "0"
	float = 0.00
	_statusDBM = StatusDBM{
		DailyEvents: 0,
		TotalEvents: 0,
		CommonEvent: "N/A",
		DailyDataRequests: 0}

	_statusSYP = StatusSYP{
		Temperature: 0.0,
		MemoryLeft: 0,
		HighestUsage: 0.0}

	_statusFH = StatusFH{
		DailyFaults: 0,
		CommonFaults: "N/A"}

	_statusNAC = StatusNAC{
		DevicesActive: 0,
		DailyBlockedDevices: 0,
		DailyUnknownDevices: 0,
		DailyAllowedDevices: 0,
		TimeEscConnected: "N/A"}

	_statusEVM = StatusEVM{
		DailyImagesTaken: 0,
		CurrentTemperature: 0.0,
		LastMotionDetected: "N/A"}

	_statusUP = StatusUP{
		LastAccessGranted: "N/A",
		LastAccessBlocked: "N/A",
		CurrentAlarmState: "OFF",
		LastUser: "N/A"}
	
}

func SetPassword(pass string) {
	password = pass
}

func failOnError(err error, msg string) {
	if err != nil {
		log.WithFields(log.Fields{
			"Message": msg, "Error": err,
		}).Trace("Rabbitmq error")
	}
}

func getTime() string {
	t := time.Now()
	log.Trace(t.Format(TIMEFORMAT))
	return t.Format(TIMEFORMAT)
}

func messages(routing_key string, value string) {
	log.Warn("Adding messages to map")
	if SubscribedMessagesMap == nil {
		log.Warn("Creation of messages map")
		SubscribedMessagesMap = make(map[uint32]*MapMessage)
		messages(routing_key, value)
	} else {
		if key_id >= 0 {
			_, valid := SubscribedMessagesMap[key_id]
			if valid {
				log.Debug("Key already exists, checking next field: ", key_id)
				key_id++
				messages(routing_key, value)
			} else {
				log.Debug("Key does not exist, adding new field: ", key_id)
				entry := MapMessage{value, routing_key, getTime(), true}
				SubscribedMessagesMap[key_id] = &entry
				key_id++
			}
		}
	}
}

func SetConnection() error {
	conn, init_err = amqp.Dial("amqp://guest:" + password + "@localhost:5672/")
	failOnError(init_err, "Failed to connect to RabbitMQ")

	ch, init_err = conn.Channel()
	failOnError(init_err, "Failed to open a channel")
	log.Trace("Beginning rabbitmq initialisation")
	log.Warn("Rabbitmq error:", init_err)
	return init_err
}

func Subscribe() {
	init := SetConnection()
	if init == nil {

		setDate()

		var topics = [6]string{
			MOTIONRESPONSE,
			STATUSSYP,
			STATUSFH,
			STATUSDBM,
			STATUSUP,
			STATUSNAC,
		}

		err := ch.ExchangeDeclare(
			EXCHANGENAME, // name
			EXCHANGETYPE, // type
			true,         // durable
			false,        // auto-deleted
			false,        // internal
			false,        // no-wait
			nil,          // arguments
		)
		failOnError(err, "EVM - Failed to declare an exchange")

		q, err := ch.QueueDeclare(
			"",    // name
			false, // durable
			false, // delete when usused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")

		for _, s := range topics {
			log.Printf("Binding queue %s to exchange %s with routing key %s",
				q.Name, EXCHANGENAME, s)
			err = ch.QueueBind(
				q.Name,       // queue name
				s,            // routing key
				EXCHANGENAME, // exchange
				false,
				nil)
			failOnError(err, "Failed to bind a queue")
		}

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto ack
			false,  // exclusive
			false,  // no local
			false,  // no wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		forever := make(chan bool)

		go func() {
			for d := range msgs {
				log.Trace("Sending message to callback")
				log.Trace(d.RoutingKey)
				s := string(d.Body[:])
				messages(d.RoutingKey, s)
				log.Debug("Checking states of received messages")
				checkState()
			}
			//This function is checked after to see if multiple errors occur then to
			//through an event message
		}()
		
		go GetWeather()

		go StatusCheck()

		log.Trace(" [*] Waiting for logs. To exit press CTRL+C")
		<-forever
	}
}

func StatusCheck() {
	done := false
	for {
		now := time.Now()
		m := now.Minute()
		if m % 4 == 0 && !done {
			PublishStatusRequest()
			done = true
		} else if m % 4 != 0 {
			done = false
		}
	}
}

func setDate() {
	_year, _month, _day = time.Now().Date()
	_messages_sent = 0
}

func checkCanSend() bool {
	year, month, day := time.Now().Date()
	if year == _year {
		if month == _month {
			if day == _day {
				if _messages_sent <= MAXMESSAGES {
					_messages_sent++
					return true
				} else {
					log.Error("Max messages sent")
					return false
				}
			} else {
				setDate()
				checkCanSend()
				_statusEVM.DailyImagesTaken = 0
			}
		}
	}
	return false
}

func PublishStatusRequest() {
	log.Debug("Publishing Status Request")
	err := ch.Publish(
		EXCHANGENAME,     // exchange
		STATUSREQUESTDBM, // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(""),
		})

	err = ch.Publish(
		EXCHANGENAME,     // exchange
		STATUSREQUESTUP, // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(""),
		})

	if err != nil {
		failOnError(err, "Failed to publish topic")
	}	
}


func PublishMotionDetected(this_time string, file string) string {
	failure := ""
	motionDetected, err := json.Marshal(&MotionDetected{
		File: file,
		Time: this_time})
	failOnError(err, "Failed to convert MotionDetected")
	log.Debug("Publishing Motion Topic")
	if err == nil {
		err = ch.Publish(
			EXCHANGENAME,     // exchange
			MOTIONDETECTED, // routing key
			false,            // mandatory
			false,            // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(motionDetected),
			})
		if err != nil {
			failOnError(err, "Failed to publish MotionDetected topic")
			failure = FAILUREPUBLISH
		}
	}
	return failure
}

func PublishFailureComponent(this_time string, this_severity int) string {
	failure := ""
	failureComponent, err := json.Marshal(&FailureMessage{
		Time:     this_time,
		Severity: this_severity})
	failOnError(err, "Failed to convert FailureMessage")

	if err == nil {
		err = ch.Publish(
			EXCHANGENAME,     // exchange
			FAILURECOMPONENT, // routing key
			false,            // mandatory
			false,            // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(failureComponent),
			})
		if err != nil {
			failOnError(err, "Failed to publish FailureComponent topic")
			failure = FAILURECOMPONENT
		}
	}
	return failure
}

func PublishEventEVM(message string, time string) string {
	failure := ""

	eventEVM, err := json.Marshal(&EventEVM{
		Component: COMPONENT,
		Message:   message,
		Time:      time})
	if err != nil {
		failure = "Failed to convert EventEVM"
	} else {
		if init_err == nil {
			log.Debug(string(eventEVM))
			err = ch.Publish(
				EXCHANGENAME, // exchange
				EVENTEVM,     // routing key
				false,        // mandatory
				false,        // immediate
				amqp.Publishing{
					ContentType: "application/json",
					Body:        []byte(eventEVM),
				})
			if err != nil {
				log.Fatal(err)
				failure = FAILUREPUBLISH
			}
		}
	}
	log.Warn(failure)
	return failure
}
