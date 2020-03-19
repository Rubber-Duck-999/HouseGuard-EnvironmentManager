package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func checkState() {
	for message_id := range SubscribedMessagesMap {
		if SubscribedMessagesMap[message_id].valid == true {
			log.Debug("Message id is: ", message_id)
			log.Debug("Message routing key is: ", SubscribedMessagesMap[message_id].routing_key)
			switch {
			case SubscribedMessagesMap[message_id].routing_key == MOTIONRESPONSE:
				log.Debug("Received a Motion Response -")
				var message MotionResponse
				json.Unmarshal([]byte(SubscribedMessagesMap[message_id].message), &message)
				if message.Severity >= 3 {
					log.Debug("Severity of 3 or more")
				}
				SubscribedMessagesMap[message_id].valid = false

			case SubscribedMessagesMap[message_id].routing_key == WEATHER:
				log.Debug("Received a Weather Topic -")
				SubscribedMessagesMap[message_id].valid = false

			default:
				log.Warn("We were not expecting this message unvalidating: ",
					SubscribedMessagesMap[message_id].routing_key)
				SubscribedMessagesMap[message_id].valid = false
			}
		}
	}

}
