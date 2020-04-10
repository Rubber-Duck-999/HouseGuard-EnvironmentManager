package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func checkState() {
	for message_id := range SubscribedMessagesMap {
		if SubscribedMessagesMap[message_id].valid == true {
			log.Trace("Message id is: ", message_id)
			log.Trace("Message routing key is: ", SubscribedMessagesMap[message_id].routing_key)
			switch {
			case SubscribedMessagesMap[message_id].routing_key == MOTIONRESPONSE:
				log.Debug("Received a Motion Response Topic")
				var message MotionResponse
				json.Unmarshal([]byte(SubscribedMessagesMap[message_id].message), &message)
				if checkCanSend() {
					if message.Severity == 2 && message.Severity == 3 {
						log.Warn("Severity of Motion from CM is mdeium")
						if motionMessage.Motion || (motionMessage.Ultrasound && motionMessage.Microwave) {
							log.Warn("Motion is apparent - notifiying service!!")
							valid := PublishMotionDetected(getTime())
							if valid != "" {
								log.Warn("Failed to publish")
							} else {
								log.Debug("Published Motion Detected Topic")
								SubscribedMessagesMap[message_id].valid = false
							}
						}
					} else if message.Severity == 4 {
						log.Debug("Severity of motion is high")
						log.Warn("Motion is apparent - notifiying service!!")
						valid := PublishMotionDetected(getTime())
						if valid != "" {
							log.Warn("Failed to publish")
						} else {
							log.Debug("Published Motion Detected Topic")
							SubscribedMessagesMap[message_id].valid = false
						}
					} else {
						log.Debug("Severity of motion is too low below 3")
					}
				} else {
					log.Error("We have received too many motions today")
				}
				SubscribedMessagesMap[message_id].valid = false
				
			default:
				log.Warn("We were not expecting this message unvalidating: ",
					SubscribedMessagesMap[message_id].routing_key)
				SubscribedMessagesMap[message_id].valid = false
			}
		}
	}

}
