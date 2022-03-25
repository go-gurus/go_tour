package main

import (
	"codecentric.de/interfaces/v2/mock_main"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_doSomething(t *testing.T) {
	ctrl := gomock.NewController(t)

	// mock the log facade in main
	logFacadeMock := mock_main.NewMocklogInterface(ctrl)

	// inject mock via mocked function
	resolveLogger = func() logInterface {
		return logFacadeMock
	}
	logFacade = resolveLogger()

	// asserts that the first calls to Info() and Error() is passed the correct strings
	// anything else will fail
	logFacadeMock.EXPECT().Info("I really dont care which logging tool is used to put this info")
	logFacadeMock.EXPECT().Error("I really dont care which logging tool is used to put this error")

	doSomething()
}
