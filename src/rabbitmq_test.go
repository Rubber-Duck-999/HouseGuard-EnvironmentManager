// rabbitmq_test.go

package main

import (
    "io/ioutil"
	"testing"
)


func TestLogicNetwork(t *testing.T) {
	value := "{ 'file': 'N/A', 'time': '2010/03/10 12:00:56', 'severity': 4 }"
	messages(MOTIONRESPONSE, value)
	checkState()
	if SubscribedMessagesMap[0].valid == true {
		t.Error("Failure")
	} else if SubscribedMessagesMap[0].routing_key == FAILURECOMPONENT {
		t.Log(SubscribedMessagesMap[0].routing_key)
		t.Error("Failure")
	}
}

func TestCleanUpOne(t *testing.T) {
	err := ioutil.WriteFile("file.jpg", []byte("Hello"), 0755)
    if err != nil {
        t.Error("Unable to write file: ", err)
	}
	cleanUp()
	if Exists("file.jpg") {
		t.Error("Failure")
	}
}

func TestCleanUpMultiple(t *testing.T) {
	err := ioutil.WriteFile("file.jpg", []byte("Hello"), 0755)
	err = ioutil.WriteFile("file2.jpg", []byte("Hello"), 0755)
	err = ioutil.WriteFile("file3.jpg", []byte("Hello"), 0755)
	err = ioutil.WriteFile("file4.jpg", []byte("Hello"), 0755)
	err = ioutil.WriteFile("file5.jpg", []byte("Hello"), 0755)
    if err != nil {
        t.Error("Unable to write file: ", err)
	}
	cleanUp()
	if Exists("file.jpg") {
		t.Error("Failure")
	}
}

func TestCheckCanSend(t *testing.T) {
	setDate()
	if checkCanSend() != true {
		t.Error("Failure")
	}
}


func TestCheckCanSendMax(t *testing.T) {
	setDate()
	for a:= 0; a < MAXMESSAGES + 5; a++ {
		value := checkCanSend()
		if value != false && a > 10 {
			t.Error("Failure")
		}
	}
}

func TestMotionDetected(t *testing.T) {
	value := "{ 'file': 'N/A', 'time': '2010/03/10 12:00:56', 'severity': 4 }"
	messages("Motion.Detected", value)
	checkState()
	if SubscribedMessagesMap[1].valid == true {
		t.Error("Failure")
	} else if SubscribedMessagesMap[1].routing_key != "Motion.Detected" {
		t.Log(SubscribedMessagesMap[1].routing_key)
		t.Error("Failure")
	}
}

func TestMotionResponseResponse(t *testing.T) {
	setDate()
	file := "../EVM.yml"
	var data ConfigTypes
	if Exists(file) {
		GetData(&data, file)
	} else {
		t.Error("Failure")
	}
	SetPassword(data.Settings.Pass)
	err := SetConnection()
	if err == nil {
		t.Log("Failure")
	}
	value := "{ 'file': 'N/A', 'time': '2010/03/10 12:00:56', 'severity': 4 }"
	messages(MOTIONRESPONSE, value)
	checkState()
	if SubscribedMessagesMap[0].valid == true {
		t.Error("Failure")
	} else if SubscribedMessagesMap[0].routing_key == FAILURECOMPONENT {
		t.Log(SubscribedMessagesMap[0].routing_key)
		t.Error("Failure")
	}
}

func TestFailure(t *testing.T) {
	setDate()
	file := "../EVM.yml"
	var data ConfigTypes
	if Exists(file) {
		GetData(&data, file)
	} else {
		t.Error("Failure")
	}
	SetPassword(data.Settings.Pass)
	err := SetConnection()
	if err == nil {
		t.Log("Failure")
	}
	PublishFailureComponent("Message", 6)
}