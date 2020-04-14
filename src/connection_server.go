package main

import (
	"bufio"
	"encoding/json"
	"net"
    "strings"

	log "github.com/sirupsen/logrus"
)

func HandleConnection() {
	log.Debug("Starting tcp connection to port 9000")
	PORT := ":9000"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		log.Debug("Listen error: ", err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Debug("Accept error: ", err)
			return
		}
		log.Trace("Serving \n", c.RemoteAddr().String())
		for {
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				log.Error("Disconnect occured :", err)
				PublishEventEVM(SENSORNETWORKDOWN, getTime())
				break
			}
	
			temp := strings.TrimSpace(string(netData))
			log.Trace("Received : ", temp)
			json.Unmarshal([]byte(temp), &motionMessage)
			if motionMessage.Motion && motionMessage.Microwave && motionMessage.Ultrasound {
				log.Warn("Motion is apparent - notifiying service!!")
				valid := PublishMotionDetected(getTime(), "N/A")
				if valid != "" {
					log.Error("Failed to publish")
				}
			}
		}
		c.Close()
	}
}
