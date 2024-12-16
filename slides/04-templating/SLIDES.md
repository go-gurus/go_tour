<!-- .slide: data-background="img/TEMPLATING/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## Templating
In this task we want to write a small program that uses the GO templating engine to create an output for 3 different unit systems.

----

### Too fast? Find source code here:
* [github.com/go-gurus/go_tour_src/tree/main/templating](https://github.com/go-gurus/go_tour_src/tree/main/templating)

----

### the task
![xss-result](img/TEMPLATING/01.jpg)<!-- .element height="60%" -->

----

### motivation

Imagine you need to present data in a well-formatted output, e.g. in `text` or in files like `html`, `json`, `xml` or `yaml`.

----

### packages
* [pkg.go.dev/text/template](https://pkg.go.dev/text/template)
* [pkg.go.dev/html/template](https://pkg.go.dev/html/template)

Both packages are based on the same basic framework and offer similar functionalities, but are specialised for different purposes.

----

### text/template
* generic templating system for processing plain text
* generation of `text` files, e.g. `YAML`, `JSON`, `INI`, `sh`, `sql` ...
* general data transformation where no special security precautions are required
* no automatic escape logic for special characters
* greater flexibility in output (no HTML-specific restrictions)

----

### html/template
* templating system developed for generating secure `HTML` output
* creation of web pages or emails with `HTML` content
* automatic escape logic, `<`, `>`, `&`, `’` converted into `HTML` entities
* recognises certain `HTML` contexts (e.g. tags, attributes, JavaScript), adapts the escaping strategy
* prevent `XSS`, malicious scripts, displayed securely in browser
* overall: protection against security vulnerabilities in dynamic content

----

### advantages

* fast and lightweight
* part of the GO standard library, no need for additional dependencies
* type safety, strict type-checking during template execution reduces runtime errors
* extensibility, you can define custom template functions to extend functionality

----

### the 1st task

In the first part of the task, we want to create a formatted text output that shows us data in three different systems of units, the metric, the imperial or US system and the maritime and aeronautical system.

----

### unit systems

<!-- | System             | Length        | Speed       |
|--------------------|---------------|-------------|
| metric             | m, km         | m/s, km/h   |
| imperial/US        | in ("), mi    | ips, mph    |
| maritime/aviation  | nm            | kn          | -->

<table>
  <thead>
    <tr>
      <th>System</th>
      <th>Length</th>
      <th>Speed</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Metric</td>
      <td>m, km</td>
      <td>m/s, km/h</td>
    </tr>
    <tr>
      <td>Imperial/US</td>
      <td>in ("), mi</td>
      <td>ips, mph</td>
    </tr>
    <tr>
      <td>Maritime/Aviation</td>
      <td>nm</td>
      <td>kn</td>
    </tr>
  </tbody>
</table>

----
<!-- .slide: data-background="img/MAIN/GOTOUR-TIME-TO-CODE-00-big.jpg" data-background-size="100%" data-background-position="50% 50%" -->

----
<!-- .slide: data-background="img/MAIN/GOTOUR-TIME-TO-CODE-00.jpg" data-background-size="100%" data-background-position="50% 50%" -->

### text/template

```golang
//main.go
package main

import (
	"os"
	"text/template"
)

type Data struct {
	Distance     float64
	DistanceUnit string
	Speed        float64
	SpeedUnit    string
	Time         float64
	TimeUnit     string
}

...

```

----

```golang
//main.go
...

func main() {
	tmpl, err := template.New("example").Parse(`
Distance: {{.Distance}} {{.DistanceUnit}}
Speed:    {{.Speed}} {{.SpeedUnit}}
You are travelling at a speed of {{.Speed}} {{.SpeedUnit}}.
You will cover a distance of {{.Distance}} {{.DistanceUnit}} in {{.Time}} {{.TimeUnit}}.
`)
	if err != nil {
		panic(err)
	}

   ...

}

```

----

```golang
//main.go
...

func main() {

   ...
   
   // metric system
	metric_data := Data{1000, "km", 100, "km/h", 10, "h"}
	err = tmpl.Execute(os.Stdout, metric_data)
	if err != nil {
		panic(err)
	}

	// imperial system
	imperial_data := Data{621.371, "mi", 62.1371, "mph", 10, "h"}
	err = tmpl.Execute(os.Stdout, imperial_data)
	if err != nil {
		panic(err)
	}
}

```

----

```golang
//main.go
...

func main() {

   ...
   
	//maritime/aviation system
	maritime_data := Data{539.957, "nm", 53.9957, "kn", 10, "h"}
	err = tmpl.Execute(os.Stdout, maritime_data)
	if err != nil {
		panic(err)
	}
}

```

----

```sh
$ go run main.go

Distance: 1000 km
Speed:    100 km/h
You are travelling at a speed of 100 km/h.
You will cover a distance of 1000 km in 10 h.

Distance: 621.371 mi
Speed:    62.1371 mph
You are travelling at a speed of 62.1371 mph.
You will cover a distance of 621.371 mi in 10 h.

Distance: 539.957 nm
Speed:    53.9957 kn
You are travelling at a speed of 53.9957 kn.
You will cover a distance of 539.957 nm in 10 h.
```

----

### the 2nd task

In the second half of the task we will try to create an `HTML` file that is exposed to an `XSS` attack. For this we will first use the package `text/template` in order to later apply and illustrate the functionality of the package `html/template`.

----
<!-- .slide: data-background="img/MAIN/GOTOUR-TIME-TO-CODE-00.jpg" data-background-size="100%" data-background-position="50% 50%" -->

Lets write a new file.

```golang
//main.go
package main

import (
	"text/template"
	"os"
)

type Data struct {
	Distance     float64
	DistanceUnit string
	Speed        float64
	SpeedUnit    string
	Time         float64
	TimeUnit     string
}
...
```

----

```golang
//main.go
...
func main() {
	tmpl, err := template.New("example").Parse(`
<!DOCTYPE html>
    <title>Travel Information</title>
</head>
<body>
    <h1>Travel Information</h1>
    <p>Distance: {{.Distance}} {{.DistanceUnit}}</p>
    <p>Speed: {{.Speed}} {{.SpeedUnit}}</p>
    <p>You are travelling at a speed of {{.Speed}} {{.SpeedUnit}}.</p>
    <p>You will cover a distance of {{.Distance}} {{.DistanceUnit}} in {{.Time}} {{.TimeUnit}}.</p>
</body>
</html>
`)
	if err != nil {
		panic(err)
	}
	...
}
```

----

* the string `<script>alert('XSS-Angriff!')</script>` is our `XSS` attack

```golang
//main.go
...
func main() {
	...
	file, err := os.Create("output.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	metricData := Data{1000, "km", 100, "<script>alert('XSS-Angriff!')</script>", 10, "h"}
	err = tmpl.Execute(file, metricData)
	if err != nil {
		panic(err)
	}
	println("HTML file 'output.html' has been created.")
}
```

----

* now let's call the go script
	```golang
	$ go run main.go
	HTML file 'output.html' has been created.
	```
* then open the `HTML` file `output.html` in the browser

----

* here you can clearly see the successful XSS attack
![xss-result](img/TEMPLATING/02.jpg)<!-- .element class="shadow-img" height="500px" -->

----

### html/template
* now let's change 1 line
	```golang
	//main.go
	package main

	import (
		"html/template"
		...
	)
	...
	```
* then open the `HTML` file `output.html` in the browser again

----

* here you can clearly see the XSS attack has no effect
![xss-result](img/TEMPLATING/03.jpg)<!-- .element class="shadow-img" height="500px" -->

----

* let's take a look at the `output.html` file
![xss-result](img/TEMPLATING/04.jpg)<!-- .element class="shadow-img" -->

----

### famous projects
* Helm, Kustomize
	```yaml
	apiVersion: apps/v1
	kind: Deployment
	metadata:
	name: {{ .Values.app.name }}
	spec:
	replicas: {{ .Values.app.replicas }}
	```
* Hugo, Staticman
	```html
	<html>
	<head>
		<title>{{ .Title }}</title>
	</head>
	<body>
		<h1>{{ .Title }}</h1>
		<p>{{ .Content }}</p>
	</body>
	</html>
	```
----

### famous projects
* Terraform/Tofu, Ansible
	```hcl
	data "template_file" "init" {
		template = <<EOT
		#!/bin/bash
		echo "Welcome {{ .Name }}"
	EOT

		vars = {
			Name = "Go Templates in Terraform"
		}
	}
	```

----

### famous projects

* Docker Compose, Swarm
	```yaml
	version: '3'
	services:
	app:
		image: {{ .Image }}
		environment:
		- ENV={{ .Env.APP_ENV }}
	```
* Prometheus, Grafana-Dashboards


----

### What we have learned
* What templates in go are suitable for.
* How to use the `text/template` package.
* How to use the `html/template` package.
* Difference between `text/template` and `html/template`.
* How to avoid `XSS` in generated `HTML` files.
* Famous projects using GO template packages.

----

### Further readings 1/3

* go package text template
  * [pkg.go.dev/text/template](https://pkg.go.dev/text/template)
* go package html template
  * [https://pkg.go.dev/html/template](https://pkg.go.dev/html/template)
* source code
  * [github.com/go-gurus/go_tour_src/tree/main/templating](https://github.com/go-gurus/go_tour_src/tree/main/templating)
* Helm
  * [helm.sh](https://helm.sh/)
  * [github.com/helm/helm](https://github.com/helm/helm)
* Kustomize
  * [github.com/kubernetes-sigs/kustomize](https://github.com/kubernetes-sigs/kustomize)

----

### Further readings 2/3

* Hugo
  * [gohugo.io](https://gohugo.io/)
  * [github.com/gohugoio/hugo](https://github.com/gohugoio/hugo)
* Staticman
  * [staticman.net](https://staticman.net/)
  * [github.com/eduardoboucas/staticman](https://github.com/eduardoboucas/staticman)
* Terraform
  * [terraform.io](https://terraform.io/)
  * [github.com/hashicorp/terraform](https://github.com/hashicorp/terraform)
* openTofu
  * [opentofu.org](https://opentofu.org/)
  * [github.com/opentofu/opentofu](https://github.com/opentofu/opentofu)


----

### Further readings 3/3

* Ansible
  * [ansible.com/](https://ansible.com/)
* Docker-Compose
  * [docs.docker.com/compose](https://docs.docker.com/compose/)
  * [github.com/docker/compose](https://github.com/docker/compose)
* Docker-Swarm
  * [docs.docker.com/engine/swarm](https://docs.docker.com/engine/swarm/)
* Prometheus
  * [prometheus.io](https://prometheus.io/)
  * [github.com/prometheus/prometheus](https://github.com/prometheus/prometheus)
* Grafana-Dashboards
  * [grafana.com/grafana/dashboards](https://grafana.com/grafana/dashboards/)

---
