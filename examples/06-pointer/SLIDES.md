<!-- .slide: data-background="img/POINTER/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## POINTER

In this task we want to create a program that initializes a memory intensive variable. In a refactoring we want to introduce pointers and clarify their advantages.

----
* first create a function to init a slice of slices

```go
// main.go
func init_image(image [][]int) [][]int {
  for key_1, val_1 := range image {
    for key_2, _ := range val_1 {
      image[key_1][key_2] = rand.Intn(256)
    }
  }
  return image
}

```

----
* now lets add a `main` method create a slice matrix and define the ranges, finally use the init method and print result

```go
// main.go
// ...
func main() {
  image := make([][]int, 1000)
  for i := 0; i < 1000; i++ {
    image[i] = make([]int, 1000)
  }

  image = init_image(image)
  fmt.Println(image)
}

```

----

### Refactoring
Now let's see how pointers improve the program and use less memory.

----

* type `*T` is a pointer to a `T` value, zero value is nil

```go
var p *int
```

* `&` operator generates a pointer to its operand

```go
i := 42
p = &i
```

* `*` operator denotes the pointer's underlying value

```go
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

----

* switch completely to pointers

```go
// main.go
package main

import (
	"fmt"
	"math/rand"
)

func init_image(image_pointer *[][]int) *[][]int {
	for key_1, val_1 := range *image_pointer {
		for key_2, _ := range val_1 {
			(*image_pointer)[key_1][key_2] = rand.Intn(256)
		}
	}
	return image_pointer
}

func main() {
	image := make([][]int, 1000)

	for i := 0; i < 1000; i++ {
		image[i] = make([]int, 1000)
	}

	pointer := &image

	pointer = init_image(pointer)
	fmt.Println(*pointer)
}

```
----

* remove now the return value

```go
// main.go
package main

import (
	"fmt"
	"math/rand"
)

func init_image(image_pointer *[][]int) {
	for key_1, val_1 := range *image_pointer {
		for key_2, _ := range val_1 {
			(*image_pointer)[key_1][key_2] = rand.Intn(256)
		}
	}
}

func main() {
	image := make([][]int, 1000)

	for i := 0; i < 1000; i++ {
		image[i] = make([]int, 1000)
	}

	init_image(&image)
	fmt.Println(image)
}

```

----

### What we have learned
* Working with range loops
* Working with slices
* Working with pointers

----

### Further readings
* Working with range loops
  * [gobyexample.com/range](https://gobyexample.com/range)
* Working with slices
  * [go.dev/tour/moretypes/7](https://go.dev/tour/moretypes/7)
  * [go.dev/doc/effective_go#slices](https://go.dev/doc/effective_go#slices)
* Working with pointers
  * [go.dev/tour/moretypes/1](https://go.dev/tour/moretypes/1)
  * [go.dev/doc/effective_go#pointers_vs_values](https://go.dev/doc/effective_go#pointers_vs_values)
  * [gobyexample.com/pointers](https://gobyexample.com/pointers)

---
