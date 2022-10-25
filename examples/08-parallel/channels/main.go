package main

import (
	"fmt"
)

func SendDataToChannel(integerChannel chan int, value int) {
	integerChannel <- value
}

func main() {
	integerChannel := make(chan int)
	go SendDataToChannel(integerChannel, 42)
	value := <-integerChannel
	fmt.Println(value)
}
