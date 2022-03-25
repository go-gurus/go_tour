// main.go
package main

import (
	"codecentric.de/interfaces/v2/golog_facade"
	"codecentric.de/interfaces/v2/logrus_facade"
	"codecentric.de/interfaces/v2/zap_facade"
	"fmt"
	"os"
	"time"
)

var logFacade logInterface = resolveLogger()

type logInterface interface {
	Info(string)
	Error(string)
}

func resolveLogger() logInterface {
	var result logInterface
	if os.Getenv("LOGGER") == "logrus" {
		result = logrus_facade.LogrusStruct{}
	} else if os.Getenv("LOGGER") == "zap" {
		result = zap_facade.NewZapStruct()
	} else if os.Getenv("LOGGER") == "golog" {
		result = golog_facade.GologStruct{}
	} else {
		fmt.Println("I don't have an idea which logger to use. So, this program will crash shortly.")
		fmt.Println("You might want to set the LOGGER envvar to prevent this. We will talk about errors later.")
	}
	return result
}

// doSomething is an example for a method in your real application having to log something
func doSomething() {
	logFacade.Info("I really dont care which logging tool is used to put this info")
	time.Sleep(time.Second)
	logFacade.Error("I really dont care which logging tool is used to put this error")
}

func main() {
	doSomething()
}
