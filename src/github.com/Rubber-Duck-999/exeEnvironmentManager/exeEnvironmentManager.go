package main

import (
	"os"

	"github.com/akamensky/argparse"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Warn("EVM - Beginning to run Environment Manager Program")
	parser := argparse.NewParser("file", "Config file for runtime purpose")
	// Create string flag
	f := parser.String("f", "config", &argparse.Options{Required: true, Help: "Necessary config"})
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		log.Error(parser.Usage(err))
		os.Exit(2)
	}

	file := *f
	var data ConfigTypes
	if Exists(file) {
		GetData(&data, file)
	} else {
		log.Error("File doesn't exist")
		os.Exit(2)
	}
	log.Trace(data.Settings.Key)
	//rabbitmq.SetSettings(data.Settings.Key,
	//	data.Settings.Api_Key)
	Subscribe()
}
