# golang_workshop

golang introduction workshop

## prerequisits
* docker

## setup
* start container
```console
$ docker run --rm -p 8000:1948 -v $(pwd)/resources:/usr/src/app containersol/reveal-md:latest
```

* check browser `http://localhost:8000`

## Further readings
* [reveal md](https://github.com/webpro/reveal-md)
* [reveal md docker file](https://hub.docker.com/r/containersol/reveal-md/)