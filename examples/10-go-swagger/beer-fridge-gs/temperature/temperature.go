package temperature

import (
	"codecentric.de/beer-fridge-go-swagger/models"
	"math/rand"
)

func GetTemperature() models.Temperature {
	min := 5
	max := 10
	return models.Temperature(rand.Intn(max-min) + min)
}
