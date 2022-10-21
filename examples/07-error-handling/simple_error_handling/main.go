// simpleErrorHandling.go
package main

import (
	"fmt"
	"os"
)

const stageEnvironmentKey = "STAGE"

func resolveService() (string, error) {
	if os.Getenv(stageEnvironmentKey) == "dev" {
		return "https://dev.fake", nil
	} else if os.Getenv(stageEnvironmentKey) == "staging" {
		return "https://stage.my.cloud", nil
	} else if os.Getenv("LOGGER") == "preprod" {
		return "https://preprod.my.cloud", nil
	} else {
		return "UNKNOWN_ENVIRONMENT", fmt.Errorf("Invalid or unknown stage")
	}
}

func main() {
	serviceUrl, err := resolveService()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Proceeding with url=%s", serviceUrl)
}
