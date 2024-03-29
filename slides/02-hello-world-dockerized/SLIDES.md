<!-- .slide: data-background="img/HELLO_WORLD_DOCKERIZED/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## Hello World Dockerized
This task is supposed to demonstrate Docker image build for Go applications.


The program should print the text `Hello World! This is Go.` to the standard output in a docker container.

----

### Too fast? Find source code here:
* [github.com/go-gurus/go_tour_src/tree/main/hello-world-dockerized](https://github.com/go-gurus/go_tour_src/tree/main/hello-world-dockerized)

----
<!-- .slide: data-background="img/HELLO_WORLD_DOCKERIZED/01.jpg" data-background-size="60%" data-background-position="50% 50%" -->

----

### Solution
<!-- .slide: data-background="img/MAIN/GOTOUR-TIME-TO-CODE-00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
```golang
// main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello World! This is Go.")
}
```

----

```Dockerfile
# Dockerfile
# build stage
FROM golang:1.17.6-alpine AS build

RUN mkdir -p /app
WORKDIR /app

# app src
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /bin/app main.go

# result stage
FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
```

----

### Building and executing Go code
* build docker image

```bash
$ docker build -t hello-image .
```
* inspect image size

```bash
$ docker images hello-image                                
REPOSITORY   TAG      IMAGE ID       CREATED         SIZE
hello-image  latest   d23a3532deaf   4 minutes ago   1.77MB
```
* run container

```bash
$ docker run --rm -it --name hello-con hello-image
Hello World! This is Go.
```

----

### What we have learned
* How to write a Go Dockerfile
* Use build stages for Go applications
* Go Docker image sizes
* faster build and deployment time
* less server/cluster workload and resource consumption

----

### Further readings

* Golang Docker file
  * [hub.docker.com/_/golang](https://hub.docker.com/_/golang)

---