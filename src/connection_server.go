package main

import (
	"bufio"
	"net"
    "strings"

	log "github.com/sirupsen/logrus"
)

func HandleConnection() {
	PORT := ":9001"
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
				log.Error(err)
				return
			}
	
			temp := strings.TrimSpace(string(netData))
			log.Debug("Received : ", temp)
		}
		c.Close()
	}
}
