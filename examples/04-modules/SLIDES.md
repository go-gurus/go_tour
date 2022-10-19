<!-- .slide: data-background="img/GolangForDevelopers-04.jpg" data-background-size="60%" data-background-position="50% 50%" -->
----

## Modules
In this task, you will have a look on modules, logging in go.

----
### logrus

Lets write a new file and import logrus.

```golang
//main.go
package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"logger": "logrus",
	}).Info("Hello from logrus")
}
```

----

### Using modules

* init modules

```bash
$ go mod init codecentric.de/hello-logrus/v2
```
* a new file `go.mod` will be generated

```golang
module codecentric.de/hello-logrus/v2
go 1.17
```

----

* install go module, a new file `go.sum` is created
```bash
$ go get github.com/sirupsen/logrus
```
* execute file `main.go`
```bash
$ go run main.go
INFO[0000] Hello from logrus     logger=logrus
```

----
### zap

* quite a higher performance compared to logrus or other logging modules

Now lets write a new file and import zap.

----

```golang
//main.go
package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Hello from zap.",
		// Structured context as strongly typed Field values.
		zap.String("logger", "zap"),
		zap.Duration("backoff", time.Second),
	)
}
```
----
* init module, install zap and run `main.go`

```bash
$ go mod init codecentric.de/hello-zap/v2
$ go get -u go.uber.org/zap
$ go run main.go
{"level":"info","ts":1644345019.7559521,"caller":"02-hello-zap/main.go:11","msg":"Hello from zap.","logger":"zap","backoff":1}
```

----

* `Dockerfile` for modules

```Dockerfile
# build stage
FROM golang:1.17.6-alpine AS build

RUN mkdir -p /app
WORKDIR /app

# build src
COPY go.mod .
COPY go.sum .
RUN go mod download

# app src
COPY .. .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app

# result stage
FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]

```

----

* what are the differences?

```Dockerfile
# Dockerfile
# ...
RUN GOOS=linux GOARCH=amd64 go build -o /bin/app main.go
# ...
```
```Dockerfile
# Dockerfile modules
# ...
# build src
COPY go.mod .
COPY go.sum .
RUN go mod download
# ...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app
# ...
```

----
* build the image

```bash
$ docker build -t hello-zap-image .
```
* run the container

```bash
$ docker run --rm -it --name hello-zap-con hello-zap-image
{"level":"info","ts":1644350565.6308103,"caller":"app/main.go:11","msg":"Hello from zap.","logger":"zap","backoff":1}
```
----

### golog

* higher performance than logrus
* lower performance than zap
* easy to use

Now lets write a new file and import golog.

----

```golang
//main.go
package main
import ( "github.com/kataras/golog" )

func main() {
	golog.SetLevel("debug")

	golog.Println("This is a raw message, no levels, no colors.")
	golog.Info("This is an info message, with colors (if the output is terminal)")
	golog.Warn("This is a warning message")
	golog.Error("This is an error message")
	golog.Debug("This is a debug message")
	golog.Fatal(`Fatal will exit no matter what,
    but it will also print the log message if logger's Level is >=FatalLevel`)
}
```
----
```bash
$ go mod init codecentric.de/hello-golog/v2
$ go get -u github.com/kataras/golog
$ go run main.go
2022/02/09 19:12 This is a raw message, no levels, no colors.
[INFO] 2022/02/09 19:12 This is an info message, with colors (if the output is terminal)
[WARN] 2022/02/09 19:12 This is a warning message
[ERRO] 2022/02/09 19:12 This is an error message
[DBUG] 2022/02/09 19:12 This is a debug message
[FTAL] 2022/02/09 19:12 Fatal will exit no matter what,
    but it will also print the log message if logger's Level is >=FatalLevel
exit status 1
```
----
### Vendoring
* What if a go module repo will be renamed, deleted or moved?
* Go programs will always compile if all module dependencies will stay accessible in the future.
* This is not always the case.
* The solution: vendoring
----
```golang
//main.go
package main
import (
	"go.uber.org/zap"
	"time"
)
func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Hello from zap.",
		zap.String("logger", "zap"),
		zap.Duration("backoff", time.Second),
	)
}
```
----
* generate vendor folder

```bash
$ go mod init codecentric.de/hello-zap-vendor
$ go get go.uber.org/zap
$ go mod vendor
```
----
* all dependencies are added

```bash
.
├── go.mod
├── go.sum
├── main.go
└── vendor
    ├── go.uber.org
    │   ├── atomic
    │   ├── multierr
    │   └── zap
    └── modules.txt

```
----
### What we have learned
* How to setup modules in go
* How to write Dockerfiles for go modules
* How to use logrus, zap and golog
* How to vendor all dependencies
----
### Further readings
* Go Modules
  * [go.dev/blog/using-go-modules](https://go.dev/blog/using-go-modules)
* Logrus
  * [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
* zap
  * [github.com/uber-go/zap](https://github.com/uber-go/zap)
* golog
  * [github.com/kataras/golog](https://github.com/kataras/golog)
* Vendoring in go
  * [go.dev/ref/mod#vendoring](https://go.dev/ref/mod#vendoring)

---
