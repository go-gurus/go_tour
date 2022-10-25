package main

import (
	"fmt"
	"math/rand"
)

func init_image(image [][]int) [][]int {
	for key_1, val_1 := range image {
		for key_2, _ := range val_1 {
			image[key_1][key_2] = rand.Intn(256)
		}
	}
	return image
}

func main() {
	image := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		image[i] = make([]int, 1000)
	}

	image = init_image(image)
	fmt.Println(image)
}
