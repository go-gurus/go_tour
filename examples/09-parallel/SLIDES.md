<!-- .slide: data-background="img/PARALLEL/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## Parallelization
In this task, we want to use go routines and channels to build a simple go load balancer.
The loadbalancer will spawn parallel go routines in the background, each of these routines will process a time intensive workload.

----

<!-- .slide: data-background="img/PARALLEL/01.jpg" data-background-size="60%" data-background-position="50% 50%" -->
----

### quick look on channels and go routines

```go
// main.go
package main

import ("fmt")

func SendDataToChannel(integerChannel chan int, value int) {
	integerChannel <- value
}

func main() {
	integerChannel := make(chan int)
	go SendDataToChannel(integerChannel, 42)
	value := <-integerChannel
	fmt.Println(value)
}
```

```go
➜ go run main.go
42

```
----

What is this program doing?
* define a function that puts values to an integer channel
* create a channel that serves integers
* executes function to put 42 to the channel in background
* get a value from channel
* print the value

----

<!-- .slide: data-background="img/PARALLEL/02.jpg" data-background-size="60%" data-background-position="50% 50%" -->
----

### simple loadbalancer using channels and go routines

----
* first start with a function that sends random values to our channel
* the function will sleep for a short time frame
* infinite loop

```go
// main.go
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
```

----

* lets create a worker function
* this function takes an integer from an input channel
* multiply the value with 1000
* sleeps a large time frame (to simulate workload)
* return a significant message
* all in an infinite loop

```golang
// main.go
// ...
func processChannel(id int, input chan int, output chan string) {
  for {
	  value, _ := <-input
	  result := value * 1000
	  time.Sleep(time.Duration(result))
	  output <- "worker: " + strconv.Itoa(id) + ", input: " + strconv.Itoa(value) + ", result: " + strconv.Itoa(result)
  }
}
```
----

* add load balancer
* create 2 channels, input (integer) and output (text)
* create subprocess to constantly add random values to input channel
* create x workers to handle the input channel and add results to output channel
* get and print values from output channel in an endless loop

----
```golang
// main.go
// ...
func main() {
    rand.Seed(time.Now().UnixNano())
    input := make(chan int)
	output := make(chan string)

    go sendDataToChannel(input)
    go processChannel(0, input, output)
    go processChannel(1, input, output)
	//...
  
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
```
----
* lets execute the load balancer (10 worker)
```shell
➜ go run main.go
worker: 8, input: 20, result: 20000
worker: 7, input: 14, result: 14000
worker: 5, input: 115, result: 115000
worker: 5, input: 16, result: 16000
worker: 3, input: 297, result: 297000
worker: 7, input: 232, result: 232000
worker: 6, input: 431, result: 431000
worker: 2, input: 719, result: 719000
worker: 0, input: 660, result: 660000
worker: 1, input: 682, result: 682000
worker: 9, input: 728, result: 728000
worker: 7, input: 389, result: 389000
worker: 8, input: 766, result: 766000
```

----

### Refactoring

----

* first lets add a loop, also change order of getting and sending data

```go
// main.go
func main() {
  // ... 
  for i := 0; i < 10; i++ {
    go processChannel(i, input, output)
  }
  go sendDataToChannel(input)
  // ...
}
```

----

* close the channels if not needed anymore

```go
// main.go
// ...
func main() {
  // ...
  input := make(chan int)
  output := make(chan string)
  defer close(input)
  defer close(output)
  // ...
}

```

----

* next, we just add a limited amount of messages to the input channel, this will break after 100 messages are handled

```go
// main.go
func sendDataToChannel(input chan int) {
  for i := 0; i < 100; i++ {
    value := rand.Intn(50)
    input <- value
    time.Sleep(time.Duration(value))
  }
}

```

----

* next, we will add non blocking channels

```go
// main.go
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

```

----

* at the end your main file should look like this

```go
// main.go
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
```

----


### What we have learned
* How to use channels
* How to close channels
* How to use non blocking channels
* How to use multiple go routines
* How to use select statement to wait on channel results
* How to use endless loops
* How to build a simple loadbalancer
----

### Further readings
* [gobyexample.com/channels](https://gobyexample.com/channels)
* [go.dev/doc/effective_go#channels](https://go.dev/doc/effective_go#channels)
* [gobyexample.com/goroutines](https://gobyexample.com/goroutines)
* [go.dev/doc/effective_go#goroutines](https://go.dev/doc/effective_go#goroutines)
* [gobyexample.com/for](https://gobyexample.com/for)
* [go.dev/doc/effective_go#for](https://go.dev/doc/effective_go#for)
* [gobyexample.com/select](https://gobyexample.com/select)

---
