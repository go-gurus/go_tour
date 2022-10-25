package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func sendDataToChannel(input chan int) {
	for {
		value := rand.Intn(50)
		input <- value
		time.Sleep(time.Duration(value))
	}
}

func processChannel(id int, input chan int, output chan string) {
	for {
		value, _ := <-input

		result := value * 1000

		time.Sleep(time.Duration(result))

		output <- "worker: " + strconv.Itoa(id) + ", input: " + strconv.Itoa(value) + ", result: " + strconv.Itoa(result)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	input := make(chan int)
	output := make(chan string)

	go sendDataToChannel(input)
	go processChannel(0, input, output)
	go processChannel(1, input, output)
	go processChannel(2, input, output)
	go processChannel(3, input, output)
	go processChannel(4, input, output)
	go processChannel(5, input, output)
	go processChannel(6, input, output)
	go processChannel(7, input, output)
	go processChannel(8, input, output)
	go processChannel(9, input, output)

	for {
		select {
		case x, ok := <-output:
			if ok {
				fmt.Println(x)
			} else {
				fmt.Println("Channel closed!")
			}
		}
	}
}
