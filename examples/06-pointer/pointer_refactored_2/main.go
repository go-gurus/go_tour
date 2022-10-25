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
