<!-- .slide: data-background="img/DEPLOYMENT/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## deployment

We want to deploy a simple go service from sources.
We use `fly.io` as a simple solution.

----

### Complete Source Code
* [github.com/go-gurus/go_tour_src/tree/main/deployment](https://github.com/go-gurus/go_tour_src/tree/main/deployment)

----

[//]: # ()
[//]: # (<!-- .slide: data-background="img/DEPLOYMENT/01.jpg" data-background-size="60%" data-background-position="50% 50%" -->)

[//]: # (----)

### fly.io
* alternative to heroku
* provides a simple and free cloud deployment for a lot of apps, services and container
* see also [fly.io](https://fly.io/)
![fly.io](img/DEPLOYMENT/02.png)<!-- .element height="400px" -->

----
* create an account
* credit card is needed, but its free
![fly.io](img/DEPLOYMENT/03.png)<!-- .element height="400px" -->
----
### flyctl
* flyctl is the fly.io command line interface
* install on MacOS

```bash
brew install flyctl
```
* install on Linux

```bash
curl -L https://fly.io/install.sh | sh
```
* install on Windows

```bash
iwr https://fly.io/install.ps1 -useb | iex
```
----
* sign up

```bash
flyctl auth signup
```

* sign in

```bash
flyctl auth login
```
----
* add a `main.go` file

```go
// main.go
package main

import (
    "embed"
    "html/template"
    "log"
    "net/http"
    "os"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := map[string]string{
            "Region": os.Getenv("FLY_REGION"),
        }

        t.ExecuteTemplate(w, "index.html.tmpl", data)
    })

    log.Println("listening on", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
```
----

* add file `templates/index.html.tmpl`

```text
<!DOCTYPE html>
<html lang="en">
  <head>
  </head>
  <body>
    <h1>Golang for Developers Workshop on Fly.io</h1>
    {{ if .Region }}
      <h2>I'm running in the {{.Region}} region</h2>
    {{end}}
  </body>
</html>
```

----
* init module

```bash
go mod init grohm.io/flyio
```

----
* launch the app

```bash
$ flyctl launch
Creating app in /Users/grohmio/repos/cc/gophers/golang-for-developers/examples/17-deployment/beer-fridge-gs
Scanning source code
Detected a Go app
Using the following build configuration:
        Builder: paketobuildpacks/builder:base
        Buildpacks: gcr.io/paketo-buildpacks/go
? Choose an app name (leave blank to generate one): flyio-grohmio

```
----

* new `fly.toml` is generated

```text
# fly.toml file generated for flyio-grohmio on 2022-12-10T00:26:40+01:00

app = "flyio-grohmio"
kill_signal = "SIGINT"
kill_timeout = 5
processes = []

[build]
  builder = "paketobuildpacks/builder:base"
  buildpacks = ["gcr.io/paketo-buildpacks/go"]

[env]
  PORT = "8080"

[experimental]
  allowed_public_ports = []
  auto_rollback = true

[[services]]
  http_checks = []
  internal_port = 8080
  processes = ["app"]
  protocol = "tcp"
  script_checks = []
  [services.concurrency]
    hard_limit = 25
    soft_limit = 20
    type = "connections"

  [[services.ports]]
    force_https = true
    handlers = ["http"]
    port = 80

  [[services.ports]]
    handlers = ["tls", "http"]
    port = 443

  [[services.tcp_checks]]
    grace_period = "1s"
    interval = "15s"
    restart_limit = 0
    timeout = "2s"

```
----

* deploy the app

```bash
$ flyctl deploy
```
----

* look at your apps in [fly.io/dashboard](https://fly.io/dashboard)

![fly.io](img/DEPLOYMENT/04.png)<!-- .element height="400px" -->

----
* look at your specific app

![fly.io](img/DEPLOYMENT/05.png)<!-- .element height="400px" -->

----

* check the link

![fly.io](img/DEPLOYMENT/06.png)<!-- .element height="400px" -->
----

### What we have learned
* Good and easy way to deploy go apps, fly.io.
* How to deploy go apps with `flyctl`


----

### Further readings
* fly.io
    * [fly.io](https://fly.io/)
    * [fly.io/docs/hands-on/install-flyctl](https://fly.io/docs/hands-on/install-flyctl/)
    * [fly.io/docs/languages-and-frameworks/golang/](https://fly.io/docs/languages-and-frameworks/golang/)
    * [fly.io/dashboard](https://fly.io/dashboard)
---