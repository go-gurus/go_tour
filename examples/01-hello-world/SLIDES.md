## Hello World
This task is supposed to demonstrate basic console I/O in Go.

The program should print the text `Hello World! This is Go.` to the standard output.

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
### What we have learned
* Go belongs to the C/Java syntactic family
* There are packages and imports, such as in Java


Useful links:
- [Golang cross compiling](https://golangcookbook.com/chapters/running/cross-compiling/0)

Note: speaker notes FTW!

---