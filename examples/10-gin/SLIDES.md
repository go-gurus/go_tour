<!-- .slide: data-background="img/GIN/00.jpg" data-background-size="60%" data-background-position="50% 50%" -->
----

## gin

Gin is a web framework written in Go (Golang).
It features a martini-like API with performance that is up to 40 times faster thanks to httprouter.
If you need performance and good productivity, you will love Gin.

----

### Getting started

* load gin modules

```bash
$ go get -u github.com/gin-gonic/gin
```

----

```go
// main.go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

----

* start service

```bash
$ go run main.go
```

* visit your browser [0.0.0.0:8080/ping](http://0.0.0.0:8080/ping)
```json
{
  message: "pong"
}
```

----

* if errors show up

```bash
...  //go:linkname must refer to declared function or variable
```

* fix
```bash
$ go get -u golang.org/x/sys
```

----

```bash
$ go mod init codecentric.de/gin-ping-pong
$ go run main.go
```
----
### A real world API Service
The Beer Fridge

- Serve temperature sensor values
- Provide information regarding the content
----
### A simple read function
in a functional style

```go
func SetupApi(r *gin.Engine, temperatureProvider func() float32) {
    api := r.Group("/api")
    {
        api.GET("/temperature", 
			composeGetTemperatureHandler(temperatureProvider))
    }
}
```
----
### A simple read function
in a functional style

```go
func getRandomTemperature() float32 {
    return 4 + rand.Float32()*3
}

func composeGetTemperatureHandler(
	    temperatureProvider func() float32) func(context *gin.Context) {
    return func(context *gin.Context) {
        context.JSON(200, gin.H{
            "temperature": temperatureProvider(),
        })
    }
}
```
----

### A simple read function
in a functional style

```bash
GET http://localhost:8080/api/temperature
```

```json
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 17 Jun 2022 13:48:16 GMT
Content-Length: 24

{
"temperature": 5.813981
}
```
----
### Evaluating Query Parameters

```go
type GetBeersFilterQuery struct {
    Origin string `form:"origin"`
}

applyQueryFilter := func(context *gin.Context, beers []Beer) []Beer {
		var query GetBeersFilterQuery

		if context.ShouldBindQuery(&query) == nil {
			// Evaluate query here
		}
		return beers
	}
```
----
### Evaluating Query Parameters
Filter method in Go 1.18 with generics
```go
if context.ShouldBindQuery(&query) == nil {
  beers = lo.Filter[Beer](beers, func(it Beer, _ int) bool {
    return query.Origin == "" || it.Origin == query.Origin
  })
}
```
----
### Mapping Result
i.E. adding HATEOAS

```go
mapHATEOAS := func(beers []Beer) []HAETEOASResource[Beer] {
		return lo.Map[Beer, HAETEOASResource[Beer]](beers, 
			func(it Beer, _ int) HAETEOASResource[Beer] {
			return HAETEOASResource[Beer]{
				Data: it,
				Links: map[string]string{
					"info":       "/" + it.urlsafePathToken(),
					"deposit":    "/" + it.urlsafePathToken() + "/deposit",
				},
			}
		})
	}
```

----

### What we have learned
* First look on gin webframework
* how to write microservices using gin
* how to build request handlers in a functional style
* how to setup a router
* how to evaluata query parameter
* how to use lodash for go

----

### Further readings
* gin
  * [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
---
