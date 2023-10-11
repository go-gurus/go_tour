<!-- .slide: data-background="img/HELLO_WORLD/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## Hello World
This task is supposed to demonstrate basic console I/O in Go.

The program should print the text `Hello World! This is Go.` to the standard output.

----

### Too fast? Find source code here:
* [github.com/go-gurus/go_tour_src/tree/main/hello-world](https://github.com/go-gurus/go_tour_src/tree/main/hello-world)

----
<!-- .slide: data-background="img/HELLO_WORLD/01.jpg" data-background-size="60%" data-background-position="50% 50%" -->

----
<!-- .slide: data-background="img/MAIN/GOTOUR-TIME-TO-CODE-00.jpg" data-background-size="100%" data-background-position="50% 50%" -->

### Solution
```golang
// main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello World! This is Go.")
}
```

----

### Building and executing Go code
Run a go file directly
```
go run main.go

Hello World! This is Go.
```

Compile and run a go file:

```
go build main.go
./main
Hello World! This is Go.
```
----

### Compile for other platforms:

```
GOOS=linux GOARCH=arm go build main.go

$ file main
main: ELF 32-bit LSB executable, ARM, EABI5 version 1 (SYSV), statically linked
```

Compile for other operating systems:
```
GOOS=windows go build main.go

$file main
main: Mach-O 64-bit executable x86_64
```

----

* e.g. build `windows/amd64` binary on `darwin/arm64`

```shell
GOOS=windows GOARCH=amd64 go build main.go
```

![windows-proof](img/HELLO_WORLD/02.png)<!-- .element height="300px" -->

----

* e.g. build `darwin/arm64` binary on `windows/amd64`

```shell
$Env:GOOS="darwin"; $Env:GOARCH="amd64"; go build main.go
```

![windows-proof](img/HELLO_WORLD/03.png)<!-- .element height="300px" -->

```shell
$ ./main
Hello World! This is Go.
```

----

### Possible Crosscompile Configurations
* [go.dev/doc/install/source#environment](https://go.dev/doc/install/source#environment)

![windows-proof](img/HELLO_WORLD/04.png)<!-- .element height="400px" -->

----
<!-- .slide: data-background="img/HELLO_WORLD/05.jpg" data-background-size="60%" data-background-position="50% 50%" -->

### Lets try a real world example
* opentofu is the open source terraform

----
<!-- .slide: data-background="img/MAIN/GOTOUR-TIME-TO-CODE-00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
* [github.com/opentofu/opentofu](https://github.com/opentofu/opentofu)
* clone the opentofu repo

```bash
$ git clone git@github.com:opentofu/opentofu.git
```

----

* now check your `GOPATH` variables, go will install packages here

```bash
$ go env GOPATH
/Users/grohmio/go
```
* add your `GOBIN` path to your `PATH`, so binaries build by go can be found, if `GOBIN`=="" default is `${GOPATH}/bin`

```
#.zshrc
...
# go binaries
export PATH="/Users/grohmio/go/bin:$PATH"
```

----

* build the binary from source, for MacOS/Darwin build with `CGO_ENABLED=1`

```bash
$ cd opentofu/
$ go build -o /Users/grohmio/go/bin/ ./cmd/tofu
go: downloading github.com/apparentlymart/go-shquot v0.0.1
go: downloading github.com/mitchellh/cli v1.1.5
go: downloading github.com/mattn/go-shellwords v1.0.4
go: downloading github.com/hashicorp/terraform-svchost v0.1.1
```

----

* check that opentofu is installed, can be used instead of terraform cli

```bash
$ tofu -h
Usage: tofu [global options] <subcommand> [args]

The available commands for execution are listed below.
The primary workflow commands are given first, followed by
less common or more advanced commands.

Main commands:
  init          Prepare your working directory for other commands
  validate      Check whether the configuration is valid
  plan          Show changes required by the current configuration
  apply         Create or update infrastructure
  destroy       Destroy previously-created infrastructure
```

----

### What we have learned
* Go belongs to the C/Java syntactic family
* There are packages and imports, such as in Java
* Hello World example
* Crosscompiling in action
* Crosscompile Configurations
* How to build binaries for huge projects
* `opentofu` tool

----

### Further readings

* Crosscompile
  * [golangcookbook.com/chapters/running/cross-compiling/0](https://golangcookbook.com/chapters/running/cross-compiling/0)
* Crosscompile Configuration
  * [go.dev/doc/install/source#environment](https://go.dev/doc/install/source#environment)
* Open Tofu
  * [github.com/opentofu/opentofu](https://github.com/opentofu/opentofu)
  * [github.com/opentofu/opentofu/blob/main/BUILDING.md](https://github.com/opentofu/opentofu/blob/main/BUILDING.md)

Note: speaker notes FTW!

---