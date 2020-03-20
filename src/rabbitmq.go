package main

import (
	"time"

	"github.com/clarketm/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel
var init_err error
var motionMessage MotionMessage

func init() {
	log.Trace("Initialised rabbitmq package")
	conn, init_err = amqp.Dial("amqp://guest:password@localhost:5672/")
	failOnError(init_err, "Failed to connect to RabbitMQ")

	ch, init_err = conn.Channel()
	failOnError(init_err, "Failed to open a channel")
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

func Subscribe() {
	log.Trace("Beginning rabbitmq initialisation")
	log.Warn("Rabbitmq error:", init_err)
	if init_err == nil {
		motionMessage.Microwave = false
		motionMessage.Ultrasound = false
		motionMessage.Motion = false

		var topics = [2]string{
			WEATHER,
			MOTIONRESPONSE,
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

		go HandleConnection()

		log.Trace(" [*] Waiting for logs. To exit press CTRL+C")
		<-forever
	}
}

func PublishMotionDetected(this_time string) string {
	failure := ""
	motionDetected, err := json.Marshal(&MotionDetected{
		Time:     this_time})
	failOnError(err, "Failed to convert MotionDetected")
	log.Debug("Publishing Motion Topic")
	if err == nil {
		err = ch.Publish(
			EXCHANGENAME, // exchange
			FAILURECOMPONENT, // routing key
			false,        // mandatory
			false,        // immediate
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
	log.Debug(string(failureComponent))

	if err == nil {
		err = ch.Publish(
			EXCHANGENAME, // exchange
			FAILURECOMPONENT, // routing key
			false,        // mandatory
			false,        // immediate
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

func PublishEventEVM(component string, message string, time string, severity int) string {
	failure := ""

	eventEVM, err := json.Marshal(&EventEVM{
		Component: component,
		Message:   message,
		Time:      time,
		Severity:  severity})
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
