package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func sendDataToChannel(input chan int) {
	for i := 0; i < 100; i++ {
		value := rand.Intn(50)
		input <- value
		time.Sleep(time.Duration(value))
	}
}

func processChannel(id int, input chan int, output chan string) {
	for {
		var result int
		select {
		case value := <-input:
			result = value * 1000
			output <- "worker: " + strconv.Itoa(id) + ", input: " + strconv.Itoa(value) + ", result: " + strconv.Itoa(result)
		default:
			// do nothing
		}
		time.Sleep(time.Duration(result))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	input := make(chan int)
	output := make(chan string)
	defer close(input)
	defer close(output)

	for i := 0; i < 10; i++ {
		go processChannel(i, input, output)
	}
	go sendDataToChannel(input)

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
