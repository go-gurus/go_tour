package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"logger": "logrus",
	}).Info("Hello from logrus")
}
