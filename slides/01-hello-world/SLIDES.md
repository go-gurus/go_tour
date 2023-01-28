<!-- .slide: data-background="img/HELLO_WORLD/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## Hello World
This task is supposed to demonstrate basic console I/O in Go.

The program should print the text `Hello World! This is Go.` to the standard output.

----

### Complete Source Code
* [github.com/go-gurus/go_tour_src/tree/main/hello-world](https://github.com/go-gurus/go_tour_src/tree/main/hello-world)

----

<!-- .slide: data-background="img/HELLO_WORLD/01.jpg" data-background-size="60%" data-background-position="50% 50%" -->
----
### Solution

```golang
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
### What we have learned
* Go belongs to the C/Java syntactic family
* There are packages and imports, such as in Java
* Hello World example
* Crosscompiling in action
* Crosscompile Configurations

----
### Further readings

* Crosscompile
  * [golangcookbook.com/chapters/running/cross-compiling/0](https://golangcookbook.com/chapters/running/cross-compiling/0)
* Crosscompile Configuration
  * [go.dev/doc/install/source#environment](https://go.dev/doc/install/source#environment)

Note: speaker notes FTW!

---