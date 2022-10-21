// simpleErrorHandling.go
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

func resolveService() (string, error) {
	// Bonus Task: Convert the return values to named values
	stageName := os.Getenv(stageEnvironmentKey)

	switch stageName {
	case "dev":
		return "https://dev.fake", nil
	case "staging":
		return "https://stage.my.cloud", nil
	case "preprod":
		return "https://preprod.my.cloud", nil
	case "":
		return invalidService, noStageNameProvided
	default:
		return invalidService, invalidStageNameProvidedError
	}
}

func remedyForMissingStageName() string {
	stdin := bufio.NewReader(os.Stdin)
	// ignore error (or any other return value) with _
	fmt.Print("Enter Service URL: ")
	input, _ := stdin.ReadString('\n')
	return input
}

func main() {
	// TODO: BONUS TASK - Retry resolve Environment via resolveService after manual input
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
