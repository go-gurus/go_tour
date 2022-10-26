# Golang Workshop for Developers

golang introduction workshop

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
$ docker run --rm -p 8000:1948 -v $(pwd)/resources:/slides webpronl/reveal-md:latest /slides --theme theme/cc.css
```

* open the slides: [http://localhost:8000](http://localhost:8000)

## Contributing

If you'd like to contribute to the project, refer to the [contributing documentation.](CONTRIBUTING.md)

## Further readings

* [reveal md](https://github.com/webpro/reveal-md)
* [reveal md docker file](https://hub.docker.com/r/containersol/reveal-md/)