# Go Tour

golang workshop, for beginners and advanced developers

## Prerequisites

* docker
* Go compiler ( >= 1.15) if you want to run the example code

## Setup Slides

* Compile slides:

```bash
./build.sh
```

* Periodic build for development purposes:

```bash
watch -n5 ./build.sh
```

* start container

```console
$ docker run --rm -p 8080:1948 -v $(pwd)/resources:/slides webpronl/reveal-md:latest /slides --theme theme/cc.css
```

* or build image and run a container

```bash
$ docker build --platform linux/amd64 --tag grohmio/go-tour:latest .
$ docker run --rm -p 8080:8080 grohmio/go-tour:latest
```

* open the slides: [http://localhost:8080](http://localhost:8080)

## Contributing

If you'd like to contribute to the project, refer to the [contributing documentation.](CONTRIBUTING.md)

## Further readings

* [reveal md](https://github.com/webpro/reveal-md)
* [reveal md docker file](https://hub.docker.com/r/containersol/reveal-md/)