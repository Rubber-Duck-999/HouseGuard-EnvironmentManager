package rabbitmq

import (
	//"encoding/json"

	log "github.com/sirupsen/logrus"
)

func messageFailure(issue bool) string {
	fail := ""
	if issue {
		fail = PublishEventEVM(COMPONENT, SERVERERROR, getTime(), SERVERSEVERITY)
	}
	return fail
}

func checkState() {
	for message_id := range SubscribedMessagesMap {
		if SubscribedMessagesMap[message_id].valid == true {
			log.Debug("Message id is: ", message_id)
			log.Debug("Message routing key is: ", SubscribedMessagesMap[message_id].routing_key)
			switch {
			case SubscribedMessagesMap[message_id].routing_key == MOTIONRESPONSE:
				log.Debug("Received a network failure message")
				//messageFailure(message.SendEmailRoutine("Serious Network failure"))
				SubscribedMessagesMap[message_id].valid = false

			case SubscribedMessagesMap[message_id].routing_key == WEATHER:
				//messageFailure(message.SendEmailRoutine("Serious Database failure"))
				//messageFailure(message.SendSMS("Serious Database failure"))
				SubscribedMessagesMap[message_id].valid = false

			default:
				log.Warn("We were not expecting this message unvalidating: ",
				SubscribedMessagesMap[message_id].routing_key)
				SubscribedMessagesMap[message_id].valid = false
			}
		}
	}

}
