<!-- .slide: data-background="img/PIPELINES/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## deployment

We want to build a simple pipeline for one of our go services.

----

### Complete Source Code
* [github.com/go-gurus/go_tour_src/tree/main/pipelines](https://github.com/go-gurus/go_tour_src/tree/main/pipelines)

----

### GitLab

* first add file `main.go`

```go
// main.go
package main

import (
	"fmt"
	"math"
)

func main() {
	var i int
	fmt.Printf("Enter a number to check: ")
	_, _ = fmt.Scanf("%d", &i)
	var result = IsPrime(i)
	fmt.Printf("result=%t ", result)
}

func IsPrime(value int) (result bool) {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

```
----

* next add testcases, add file `main_test.go`

```go
// main_test.go
package main

import (
	"testing"
)

func TestPrimeCheckerTheNaiveWay(t *testing.T) {
	t.Run("should return FALSE when no prime number given", func(t *testing.T) {
		if IsPrime(4) == true {
			t.Fatal("Reported IsPrime=true for 4")
		}
	})

	t.Run("should return TRUE when prime number given", func(t *testing.T) {
		if IsPrime(7) == false {
			t.Fatal("Reported IsPrime=true for 7")
		}
	})
}

func TestPrimeCheckerTableDriven(t *testing.T) {
	cases := []struct {
		input          int
		expectedResult bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{7, true},
	}

	for _, e := range cases {
		t.Run("Should return expected result ", func(t *testing.T) {
			result := IsPrime(e.input)
			if result != e.expectedResult {
				t.Fatalf("Unexpected Result input=%d expected=%t actual=%t", e.input, e.expectedResult, result)
			}
		})
	}
}

```

----

* initialize module

```bash
go mod init grohm.io/pipelines_showcase
```

* create empty file `go.sum`

----

* create pipeline file `.gitlab-ci.yml`

```yaml
# .gitlab-ci.yml
image: golang:1.19-alpine
```

----

* next add stages for our pipeline

```yaml
# .gitlab-ci.yml
# ...
stages:
  - lint
  - test
  - build
```

----

* add the fmt job

```yaml
# .gitlab-ci.yml
# ...
fmt:
  stage: lint
  script:
    - go fmt
    - go vet
```

----

* add the test job, get the coverage report artifact

```yaml
# .gitlab-ci.yml
# ...
test:
  stage: test
  script:
    - go test -coverprofile=coverage.out
    - go tool cover -html=coverage.out -o coverage.html
  artifacts:
    paths:
      - coverage.html
    expire_in: 1 week
```

----

* add the build job, get the binary artifact

```yaml
# .gitlab-ci.yml
# ...
build:
  stage: build
  script:
    - go build
  artifacts:
    paths:
      - pipelines_showcase
    expire_in: 1 week
```

----

* now lets add the docker build

```dockerfile
# Dockerfile
# build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /app
WORKDIR /app

# build src
COPY go.mod .
COPY go.sum .
RUN go mod download

# app src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app

# result stage
FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
```

----

```yaml
# .gitlab-ci.yml
# ...
docker-build:
  stage: containerize
  image:
    name: gcr.io/kaniko-project/executor:debug
    entrypoint: [ "" ]
  variables:
    KUBERNETES_CPU_REQUEST: 1
    KUBERNETES_CPU_LIMIT: 1
    KUBERNETES_MEMORY_REQUEST: 2048Mi
    KUBERNETES_MEMORY_LIMIT: 2048Mi
  needs: ["build"]
  script:
    - mkdir -p /kaniko/.docker
    - echo "{\"auths\":{\"$CI_REGISTRY\":{\"username\":\"$CI_REGISTRY_USER\",\"password\":\"$CI_REGISTRY_PASSWORD\"}}}" > /kaniko/.docker/config.json
    - /kaniko/executor --context $CI_PROJECT_DIR --dockerfile $CI_PROJECT_DIR/Dockerfile --destination $CI_REGISTRY_IMAGE:$CI_PIPELINE_IID
```

----

* to be complete, lets add a release tag process

```yaml
# .gitlab-ci.yml
# ...
release_job:
  stage: release_tagging
  image: registry.gitlab.com/gitlab-org/release-cli:latest
  needs:
    - job: docker-build
      artifacts: true
  script:
    - echo "running release_job for $CI_PIPELINE_IID"
  release:
    name: 'Release $CI_PIPELINE_IID'
    description: 'Created using the release-cli'
    tag_name: '$CI_PIPELINE_IID'
    ref: '$CI_COMMIT_SHA'
```
----

* lets add stages and avoid running on tag pipelines

```yaml
# .gitlab-ci.yml
# ...
stages:
  - lint
  - test
  - build
  - containerize
  - release_tagging

workflow:
  rules:
    - if: $CI_COMMIT_TAG
      when: never
    - when: always
# ...
```
----

* add changes, commit and push
* the pipeline should look like this

![fly.io](img/PIPELINES/01.png)<!-- .element height="400px" -->

----

* enter the test job, check the coverage, check artifacts

![fly.io](img/PIPELINES/02.png)<!-- .element height="400px" -->

----

* check coverage report artifact

![fly.io](img/PIPELINES/03.png)<!-- .element height="400px" -->

----

* report should look like this

![fly.io](img/PIPELINES/04.png)<!-- .element height="400px" -->

----

* check also the docker image in the registry

![fly.io](img/PIPELINES/05.png)<!-- .element height="400px" -->

----

### What we have learned

* How to build a GitLab pipeline for go services

----

### Further readings
* GitLab Pipelines
    * [docs.gitlab.com/ee/ci/pipelines](https://docs.gitlab.com/ee/ci/pipelines/)
* Golang GitLab Template
    * [gitlab.com/gitlab-org/gitlab/-/blob/master/lib/gitlab/ci/templates/Go.gitlab-ci.yml](https://gitlab.com/gitlab-org/gitlab/-/blob/master/lib/gitlab/ci/templates/Go.gitlab-ci.yml)
---
