# How to contribute

## Changing the Slides

* start container

```console
$ docker run --rm -p 8000:1948 -v $(pwd)/resources:/slides webpronl/reveal-md:latest /slides --theme theme/cc.css
```

* run periodic slides update

```console
  $ watch -n5 ./build.sh  
```

## Terraform

* init project
```shell
cd devops
terraform init \
-backend-config="address=https://gitlab.codecentric.de/api/v4/projects/5409/terraform/state/main" \
-backend-config=username="..." \
-backend-config=password="..."
```

## Further readings

* [reveal md](https://github.com/webpro/reveal-md)
* [reveal md docker file](https://hub.docker.com/r/containersol/reveal-md/)