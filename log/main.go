package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"animal": "cat",
		"size":   20,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"animal": "bird",
		"size":   30,
	}).Error("The ice breaks!")

	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Warn("I'll be logged with common and other field")
	contextLogger.Info("Cek")
	contextLogger.WithFields(log.Fields{
		"data": "response",
	}).Info("Cek")
}
