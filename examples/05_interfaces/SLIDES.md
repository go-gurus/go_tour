## Interface
In this task, we want to use interfaces to build a go service that is able to use logrus, zap and golog as logger.
The service is configured via environment variable `LOGGER=[logrus|zap|golog]`

----
### Create an interface

```golang
// main.go
package main

type logInterface interface {
	Info(string)
	Error(string)
}
// ...
```

----

* now lets create some folders and files (a.k.a. packages)
```bash
golog_facade/
    golog_facade.go
logrus_facade/
    logrus_facade.go
zap_facade/
    zap_facade.go
main.go
```

----

* lets create the `logrus` logging fassade
* notice: the struct `LogrusStruct` has the two methods that are needed to fulfill the interface `logInterface`

----

```golang
// logrus_facade/logrus_facade.go
package logrus_fassade
import ("github.com/sirupsen/logrus")

type LogrusStruct struct{}
func (logger LogrusStruct) Info(msg string) { 
	logrus.Debug(msg + " (logrus)") 
}

func (logger LogrusStruct) Error(msg string) {
	logrus.Error(msg + " (logrus)") 
}
```

----

* lets write the `golog` logging fassade

```golang
// golog_facade/golog_facade.go
package golog_fassade
import ("github.com/kataras/golog")

type GologStruct struct{}
func (s GologStruct) Info(msg string) { 
	golog.Debug(msg + " (golog)") 
}

func (s GologStruct) Error(msg string) { 
	golog.Error(msg + " (golog)") 
}
```

----

* lets write the `zap` logging fassade

```golang
// zap_facade/zap_facade.go
package zap_fassade
import ("go.uber.org/zap")

type ZapStruct struct { logger zap.Logger}

func (s ZapStruct) Info(msg string) {
	defer s.logger.Sync()
	s.logger.Info(msg + " (zap)")
}
func (s ZapStruct) Error(msg string) {
	defer s.logger.Sync()
	s.logger.Debug(msg + " (zap)")
}
func NewZapStruct() ZapStruct {
	logger, _ := zap.NewProduction()
	result := ZapStruct{*logger}
	return result
}
```

----

* add a log resolver

```golang
// main.go
// ...
func resolveLogger() logInterface
    var result logInterface
    if os.Getenv("LOGGER") == "logrus" {
        result = logrus_fassade.LogrusStruct{}
    } else if os.Getenv("LOGGER") == "zap" {
        result = zap_fassade.NewZapStruct()
    } else if os.Getenv("LOGGER") == "golog" {
        result = golog_fassade.GologStruct{}
    } else {
        fmt.Println("Unknown logger, please set $LOGGER envvar.")
    }
    return result
}
```

----

* add a example function

```golang
// main.go
// ...
var logFacade logInterface = resolveLogger()

func doSomething() {
    logFacade.Info("I really dont care which logging tool is used to put this info")
    time.Sleep(time.Second)
    logFacade.Error("I really dont care which logging tool is used to put this error")
}
func main() { doSomething() }
```

----

* init project and download modules

```bash
$ go mod init codecentric.de/interfaces/v2
$ go get github.com/sirupsen/logrus
$ go get -u go.uber.org/zap
$ go get -u github.com/kataras/golog
```

----

* execute with different configurations

```bash
$ LOGGER=golog go run main.go
[INFO] 2022/02/12 20:20 i really dont care which logging tool is used to put this info (golog)
[ERRO] 2022/02/12 20:20 i really dont care which logging tool is used to put this error (golog)
```

```bash
$ LOGGER=logrus go run main.go
INFO[0000] i really dont care which logging tool is used to put this info (logrus) 
ERRO[0001] i really dont care which logging tool is used to put this error (logrus)
```

```bash
$ LOGGER=zap go run main.go
{"level":"info","ts":1644693743.249532,"caller":"zap_fassade/zap_fassade.go:14","msg":"i really dont care which logging tool is used to put this info (zap)"}
{"level":"error","ts":1644693744.2508721,"caller":"zap_fassade/zap_fassade.go:19","msg":"i really dont care which logging tool is used to put this error (zap)",
"stacktrace":"codecentric.de/interfaces/v2/zap_fassade.ZapStruct.Error\n\t/Users/grohmio/repos/cc/gitlab/golang_workshop/examples/05_interfaces/zap_fassade/zap_fassade.go:19\nmain.doSomething\n\t/Users/grohmio/repos/cc/gitlab/golang_workshop/examples/05_interfaces/main.go:35\nmain.main\n\t/Users/grohmio/repos/cc/gitlab/golang_workshop/examples/05_interfaces/main.go:39\nruntime.main\n\t/usr/local/opt/go/libexec/src/runtime/proc.go:255"}
```

----

### What we have learned
* How to interfaces
* How to use subpackages
* How to use mocking

----

### Further readings
* [TODO](TODO)

---
