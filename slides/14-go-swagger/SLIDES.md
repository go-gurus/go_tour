<!-- .slide: data-background="img/GO_SWAGGER/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## go-swagger

We want to build the `beer-fridge` service again.
This service using the interface first approach with `go-swagger` code generation.

----

### Too fast? Find source code here:
* [github.com/go-gurus/go_tour_src/tree/main/go-swagger](https://github.com/go-gurus/go_tour_src/tree/main/go-swagger)

----
<!-- .slide: data-background="img/GO_SWAGGER/01.jpg" data-background-size="60%" data-background-position="50% 50%" -->

----
<!-- .slide: data-background="img/MAIN/GOTOUR-TIME-TO-CODE-00.jpg" data-background-size="100%" data-background-position="50% 50%" -->

### get go-swagger
* install `go-swagger` via docker on mac and linux

```shell
docker pull quay.io/goswagger/swagger
alias swagger='docker run --rm -it  --user $(id -u):$(id -g) -e \
GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'
swagger version
```

* install `go-swagger` via docker on windows

```shell
docker run --rm -it --env GOPATH=/go -v %CD%:/go/src -w /go/src quay.io/goswagger/swagger
```

* check swagger version

```shell
$ swagger version
version: v0.29.0
commit: 53696caa1e8a4e5b483c87895d54eda202beb3b0
```

----

* init a new project

```shell
swagger init spec \
  --title "A beer fridge service" \
  --description "Beer fridge service build with go-swagger" \
  --version 1.0.0 \
  --scheme http \
  --consumes application/io.grohm.go-workshop.beer-fridge.v1+json \
  --produces application/io.grohm.go-workshop.beer-fridge.v1+json
```

----

* check the result, file `swagger.yml`

```yaml
# swagger.yml
consumes:
- application/io.grohm.go-workshop.beer-fridge.v1+json
info:
  description: Beer fridge service build with go-swagger
  title: A beer fridge service
  version: 1.0.0
paths: {}
produces:
- application/io.grohm.go-workshop.beer-fridge.v1+json
schemes:
- http
swagger: "2.0"
```

----

* validate the interface file

```shell
➜ swagger validate swagger.yml 
2022/07/01 19:57:02 
The swagger spec at "swagger.yml" is valid against swagger specification 2.0
2022/07/01 19:57:02 
The swagger spec at "swagger.yml" showed up some valid but possibly unwanted constructs.
2022/07/01 19:57:02 See warnings below:
2022/07/01 19:57:02 - WARNING: spec has no valid path defined

```

----

* lets adapt the interface to our needs
* add definitions for `beer` and `teperature` of our fridge

```yaml
# swagger.yml
# ...
definitions:
  beer:
    type: object
    required:
      - title
      - origin
      - volume-percentage
    properties:
      id:
        type: integer
        format: int64
        readOnly: true
      title:
        type: string
        minLength: 1
      origin:
        type: string
        minLength: 1
      volume-percentage:
        type: number
        format: float
        minLength: 1
  temperature:
    type: integer
    format: int64
    readOnly: true
```

----

* add endpoints for `/beers` and `/temperature`

```yaml
paths:
  /beers:
    get:
      tags:
        - beers
      parameters:
        - name: limit
          in: query
          type: integer
          format: int32
          default: 10
      responses:
        200:
          description: list the beer operations
          schema:
            type: array
            items:
              $ref: "#/definitions/beer"
  /temperature:
    get:
      tags:
        - fridge
      responses:
        200:
          description: return the current fridge temperature
          schema:
            $ref: "#/definitions/temperature"
```

----

* init module, generate service and get modules

```shell
➜ go mod init grohm.io/beer-fridge-go-swagger
➜ swagger generate server -A beer-fridge -f ./swagger.yml
➜ go get -u -f ./...
```

----

* check generated code

```shell
➜ tree 
.
├── cmd
│ └── beer-fridge-server
│     └── main.go
├── go.mod
├── go.sum
├── models
│ ├── beer.go
│ └── temperature.go
├── restapi
│ ├── configure_beer_fridge.go
│ ├── doc.go
│ ├── embedded_spec.go
│ ├── operations
│ │ ├── beer_fridge_api.go
│ │ ├── beers
│ │ │ ├── get_beers.go
│ │ │ ├── get_beers_parameters.go
│ │ │ ├── get_beers_responses.go
│ │ │ └── get_beers_urlbuilder.go
│ │ └── fridge
│ │     ├── get_temperature.go
│ │     ├── get_temperature_parameters.go
│ │     ├── get_temperature_responses.go
│ │     └── get_temperature_urlbuilder.go
│ └── server.go
└── swagger.yml

7 directories, 19 files
```

----

* start service

```shell
➜ go run cmd/beer-fridge-server/main.go
2022/08/28 16:34:16 Serving beer fridge at http://127.0.0.1:54746
```

* visit service under [127.0.0.1:54746/beers](http://127.0.0.1:65480/beers) (take the custom port into account)

```html
"operation beers.GetBeers has not yet been implemented"
```

----

* build the server binary

```shell
➜ go build -o beer-fridge-server ./cmd/beer-fridge-server/main.go
➜ ./beer-fridge-server --help
```

* run the server as binary

```shell
➜ ./beer-fridge-server --port 8080
2022/08/05 19:18:18 Serving beer fridge at http://127.0.0.1:8080
```

* visit service under [127.0.0.1:8080/beers](http://127.0.0.1:8080/beers)

```html
"operation beers.GetBeers has not yet been implemented"
```

----

* add `error` definition
* enlarge endpoint `/beers`, add beer

```yaml
# swagger.yml
definitions:
  # ...
  error:
    type: object
    required:
      - message
    properties:
      code:
        type: integer
        format: int64
      message:
        type: string
# ...
paths:
  /beers:
    #...
    post:
      tags:
        - todos
      operationId: addOne
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/beer"
      responses:
        201:
          description: Created
          schema:
            $ref: "#/definitions/beer"
        default:
          description: error
          schema:
            $ref: "#/definitions/error" 
```

----

* enlarge endpoint `/beers`, delete a beer

```yaml
# swagger.yml
/paths:
  beers:
    delete:
      tags:
        - beers
      operationId: destroyOne
      responses:
        204:
          description: Deleted
        default:
          description: error
          schema:
            $ref: "#/definitions/error"
```

----

* final file should look like this

```yaml
consumes:
- application/io.grohm.go-workshop.beer-fridge.v1+json
info:
  description: Beer fridge service build with go-swagger
  title: A beer fridge service
  version: 1.0.0
produces:
- application/io.grohm.go-workshop.beer-fridge.v1+json
schemes:
- http
swagger: "2.0"
definitions:
  beer:
    type: object
    required:
      - title
      - origin
      - volume-percentage
    properties:
      id:
        type: integer
        format: int64
        readOnly: true
      title:
        type: string
        minLength: 1
      origin:
        type: string
        minLength: 1
      volume-percentage:
        type: number
        format: float
        minLength: 1
  temperature:
    type: integer
    format: int64
    readOnly: true
  error:
    type: object
    required:
      - message
    properties:
      code:
        type: integer
        format: int64
      message:
        type: string
paths:
  /beers:
    get:
      operationId: getAllBeers
      tags:
        - beers
      parameters:
        - name: limit
          in: query
          type: integer
          format: int32
          default: 20
      responses:
        200:
          description: list the beer operations
          schema:
            type: array
            items:
              $ref: "#/definitions/beer"
    post:
      tags:
        - beers
      operationId: addOne
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/beer"
      responses:
        201:
          description: Created
          schema:
            $ref: "#/definitions/beer"
        default:
          description: error
          schema:
            $ref: "#/definitions/error"
  /beers/{id}:
    delete:
      tags:
        - beers
      operationId: destroyOne
      parameters:
        - type: integer
          format: int64
          name: id
          in: path
          required: true
      responses:
        204:
          description: Deleted
        default:
          description: error
          schema:
            $ref: "#/definitions/error"
  /temperature:
    get:
      operationId: getTemperature
      tags:
        - fridge
      responses:
        200:
          description: return the current fridge temperature
          schema:
            $ref: "#/definitions/temperature"
```

----

* build the server again

```shell
swagger generate server -A beer-fridge -f ./swagger.yml
go get -u -f ./...
```

* lets add a simple beer data container

```go
// beer_container/beer_container.go
package beers

import (
	"grohm.io/beer-fridge-go-swagger/models"
	"github.com/go-openapi/errors"
	"sync"
	"sync/atomic"
)

var beerList = make(map[int64]*models.Beer)
var lastID int64
var beerListLock = &sync.Mutex{}

func newBeerID() int64 {
	return atomic.AddInt64(&lastID, 1)
}
```

----

* add a beer to list

```go
// beer_container/beer_container.go
//...
func AddBeer(beer *models.Beer) error {
	if beer == nil {
		return errors.New(500, "beer must be present")
	}

	beerListLock.Lock()
	defer beerListLock.Unlock()

	newID := newBeerID()
	beer.ID = newID
	beerList[newID] = beer

	return nil
}
```

----

* drink and remove a beer from list

```go
// beer_container/beer_container.go
//...
func DeleteBeer(id int64) error {
	beerListLock.Lock()
	defer beerListLock.Unlock()

	_, exists := beerList[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	delete(beerList, id)
	return nil
}
```

----

* show all beer's on list

```go
// beer_container/beer_container.go
//...
func AllBeers(limit int32) (result []*models.Beer) {
	result = make([]*models.Beer, 0)
	for _, beer := range beerList {
		if len(result) >= int(limit) {
			return
		}
		result = append(result, beer)
	}
	return
}
```

----

* now lets integrate the beer container into the api, add following lines

```go
// restapi/configure_beer_fridge.go
// ...
func configureAPI(api *operations.BeerFridgeAPI) http.Handler {
	// ...
	api.BeersAddOneHandler = beers.AddOneHandlerFunc(func(params beers.AddOneParams) middleware.Responder {
	    if err := beer_container.AddBeer(params.Body); err != nil {
	        return beers.NewAddOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
	    }
	    return beers.NewAddOneCreated().WithPayload(params.Body)
	})

	api.BeersDestroyOneHandler = beers.DestroyOneHandlerFunc(func(params beers.DestroyOneParams) middleware.Responder {
	    beer_container.DeleteBeer(params.ID)
	    return beers.NewDestroyOneNoContent()
	})

	api.BeersGetAllBeersHandler = beers.GetAllBeersHandlerFunc(func(params beers.GetAllBeersParams) middleware.Responder {
	    mergedParams := beers.NewGetAllBeersParams()
	    if params.Limit != nil {
	        mergedParams.Limit = params.Limit
	    }
	    return beers.NewGetAllBeersOK().WithPayload(beer_container.AllBeers(*mergedParams.Limit))
	})
}
```

----

* dont forget to add the endpoint `temperature`

```go
// temperature/temperature.go
package temperature

import (
	"grohm.io/beer-fridge-go-swagger/models"
	"math/rand"
)

func GetTemperature() models.Temperature {
	min := 5
	max := 10
	return models.Temperature(rand.Intn(max-min) + min)
}

```

----

* also integrate the temperature into the api, add following lines

```go
// restapi/configure_beer_fridge.go
// ...
func configureAPI(api *operations.BeerFridgeAPI) http.Handler {
	// ...
	api.FridgeGetTemperatureHandler = fridge.GetTemperatureHandlerFunc(func(params fridge.GetTemperatureParams) middleware.Responder {
        return fridge.NewGetTemperatureOK().WithPayload(temperature.GetTemperature())
	})
}
```

----

* rebuild and start the server

```shell
go build -o beer-fridge-server ./cmd/beer-fridge-server/main.go
./beer-fridge-server --port 8080
```

* add 3 beers

```shell
curl -i localhost:8080/beers -d "{\"title\":\"Three Floyds Brewing Co.\", \"origin\":\"Munster, Ind.\", \"volume-percentage\": 5}" -H 'Content-Type: application/io.grohm.go-workshop.beer-fridge.v1+json'
curl -i localhost:8080/beers -d "{\"title\":\"The Alchemist Heady Topper\", \"origin\":\"Waterbury, Vt.\", \"volume-percentage\": 6}" -H 'Content-Type: application/io.grohm.go-workshop.beer-fridge.v1+json'
curl -i localhost:8080/beers -d "{\"title\":\"Founders KBS (Kentucky Breakfast Stout)\", \"origin\":\"Grand Rapids, Mich.\", \"volume-percentage\": 7}" -H 'Content-Type: application/io.grohm.go-workshop.beer-fridge.v1+json'
```

----

* get all beers

```shell
curl -i localhost:8080/beers
```

* get the temperature

```shell
curl -i localhost:8080/temperature
```

----

* delete a beer

```shell
curl -i localhost:8080/beers/1 -X DELETE -H 'Content-Type: application/io.grohm.go-workshop.beer-fridge.v1+json'
```

* lets try something

```shell
$ curl -i localhost:8080/beers/3 -H 'Content-Type: application/io.grohm.go-workshop.beer-fridge.v1+json'
HTTP/1.1 405 Method Not Allowed
Allow: DELETE
Content-Type: application/json
Date: Sat, 27 Aug 2022 17:56:28 GMT
Content-Length: 68

{"code":405,"message":"method GET is not allowed, but [DELETE] are"}%  
```

----

* specific Dockerfile for this service

```dockerfile
# build stage
FROM golang:1.17.6-alpine AS build

RUN mkdir -p /app
WORKDIR /app

# build src
COPY go.mod .
COPY go.sum .
RUN go mod download

# app src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o beer-fridge-server ./cmd/beer-fridge-server/main.go

# result stage
FROM scratch
COPY --from=build /app/beer-fridge-server /beer-fridge-server
EXPOSE 8080
EXPOSE 443
EXPOSE 80
ENTRYPOINT ["/beer-fridge-server", "--port", "8080"]
```

----

### What we have learned
* How to write an open api yml file.
* How to install and use go-swagger.
* How to build an REST API from ground up.
* How to develop with interface first approach.

----

### Further readings
* go-swagger
    * [goswagger.io/](https://goswagger.io/)
    * [github.com/go-swagger/go-swagger](https://github.com/go-swagger/go-swagger)
* install go-swagger
    * [goswagger.io/install.html](https://goswagger.io/install.html)
* go-swagger tutorial
    * [goswagger.io/tutorial/todo-list.html](https://goswagger.io/tutorial/todo-list.html)
* go-swagger complete example
    * [github.com/go-swagger/go-swagger/tree/master/examples/tutorials/todo-list/server-complete](https://github.com/go-swagger/go-swagger/tree/master/examples/tutorials/todo-list/server-complete)

---