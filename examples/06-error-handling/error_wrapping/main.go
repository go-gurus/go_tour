package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const stageEnvironmentKey = "STAGE"
const invalidService = "UNKNOWN_ENVIRONMENT"

var noStageNameProvided = errors.New("No Stage name provided")
var invalidStageNameProvidedError = errors.New("Invalid stage name provided")

func resolveService() (serviceUrl string, err error) {
	stageName := os.Getenv(stageEnvironmentKey)

	serviceUrl = invalidService
	switch stageName {
	case "dev":
		serviceUrl = "https://dev.fake"
	case "staging":
		serviceUrl = "https://stage.my.cloud"
	case "preprod":
		serviceUrl = "https://preprod.my.cloud"
	case "":
		err = noStageNameProvided
	default:
		err = fmt.Errorf("%w . %s is not a known stageName", invalidStageNameProvidedError, stageName)
	}
	return
}

func remedyForMissingStageName() string {
	stdin := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Service URL: ")
	input, _ := stdin.ReadString('\n')
	return input
}

func main() {
	serviceUrl, err := resolveService()
	if err != nil {
		switch {
		case errors.Is(err, noStageNameProvided):
			fmt.Println("No stage name provided.")
			serviceUrl = remedyForMissingStageName()

		case errors.Is(err, invalidStageNameProvidedError):
			panic(err)
		}
	}
	fmt.Printf("Proceeding with url=%s", serviceUrl)
}
