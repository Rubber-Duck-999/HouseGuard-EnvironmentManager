package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func cleanUp() {

	dirname := "." + string(filepath.Separator)

	d, err := os.Open(dirname)
	if err != nil {
		log.Warn(err)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		log.Warn(err)
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".jpg" {
				os.Remove(file.Name())
				log.Warn("Deleted ", file.Name())
			}
		}
	}
}

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
					log.Debug("Severity of motion is high")
					log.Warn("Motion is apparent - notifiying service!!")
					valid := PublishMotionDetected(getTime(), message.File)
					if message.File == "N/A" {
						driveAddFile(message.File)
					}
					if valid != "" {
						log.Warn("Failed to publish")
					} else {
						log.Debug("Published Motion Detected Topic")
						SubscribedMessagesMap[message_id].valid = false
					}
				} else {
					log.Error("We have received too many motions today")
				}
				SubscribedMessagesMap[message_id].valid = false
				cleanUp()

			case SubscribedMessagesMap[message_id].routing_key == STATUSDBM:
				log.Debug("Received a Status DBM Topic")
				var message StatusDBM
				json.Unmarshal([]byte(SubscribedMessagesMap[message_id].message), &message)
				_statusDBM = message
				driveUpdateStatus()
				SubscribedMessagesMap[message_id].valid = false

			case SubscribedMessagesMap[message_id].routing_key == STATUSSYP:
				log.Debug("Received a Status SYP Topic")
				var message StatusSYP
				json.Unmarshal([]byte(SubscribedMessagesMap[message_id].message), &message)
				_statusSYP = message
				driveUpdateStatus()
				SubscribedMessagesMap[message_id].valid = false

			case SubscribedMessagesMap[message_id].routing_key == STATUSFH:
				log.Debug("Received a Status FH Topic")
				var message StatusFH
				json.Unmarshal([]byte(SubscribedMessagesMap[message_id].message), &message)
				_statusFH = message
				driveUpdateStatus()
				SubscribedMessagesMap[message_id].valid = false

			case SubscribedMessagesMap[message_id].routing_key == STATUSNAC:
				log.Debug("Received a Status NAC Topic")
				var message StatusNAC
				json.Unmarshal([]byte(SubscribedMessagesMap[message_id].message), &message)
				_statusNAC = message
				driveUpdateStatus()
				SubscribedMessagesMap[message_id].valid = false

			case SubscribedMessagesMap[message_id].routing_key == STATUSUP:
				log.Debug("Received a Status UP Topic")
				var message StatusUP
				json.Unmarshal([]byte(SubscribedMessagesMap[message_id].message), &message)
				_statusUP = message
				driveUpdateStatus()
				SubscribedMessagesMap[message_id].valid = false

			default:
				log.Warn("We were not expecting this message unvalidating: ",
					SubscribedMessagesMap[message_id].routing_key)
				SubscribedMessagesMap[message_id].valid = false
			}
		}
	}

}
