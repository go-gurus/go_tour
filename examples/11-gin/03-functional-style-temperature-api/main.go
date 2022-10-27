package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"math/rand"
	"net/url"
)

func getRandomTemperature() float32 {
	return 4 + rand.Float32()*3
}

// => Function Composition
func composeGetTemperatureHandler(temperatureProvider func() float32) func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"temperature": temperatureProvider(),
		})
	}
}

type Beer struct {
	Title            string
	Origin           string
	VolumePercentage float32
}

type HAETEOASResource[T any] struct {
	Data  T
	Links map[string]string
}

func getInitialBeerList() []Beer {
	return []Beer{
		{
			Title:            "Oostmalle Trappist Dubbel",
			Origin:           "BE",
			VolumePercentage: 8.7,
		},
		{
			Title:            "Solingen 7am Escalation IPA",
			Origin:           "DE",
			VolumePercentage: 7.0,
		},
	}
}

func (b *Beer) urlsafePathToken() string {
	return url.QueryEscape(b.Title)
}

func composeGetAllBeersHandler() func(context *gin.Context) {
	type GetBeersFilterQuery struct {
		Origin string `form:"origin"`
	}

	applyQueryFilter := func(context *gin.Context, beers []Beer) []Beer {
		var query GetBeersFilterQuery

		if context.ShouldBindQuery(&query) == nil {
			beers = lo.Filter[Beer](beers, func(it Beer, _ int) bool {
				return query.Origin == "" || it.Origin == query.Origin
			})
		}
		return beers
	}

	mapHATEOAS := func(beers []Beer) []HAETEOASResource[Beer] {
		return lo.Map[Beer, HAETEOASResource[Beer]](beers, func(it Beer, _ int) HAETEOASResource[Beer] {
			return HAETEOASResource[Beer]{
				Data: it,
				Links: map[string]string{
					"info":       "/" + it.urlsafePathToken(),
					"deposit":    "/" + it.urlsafePathToken() + "/deposit",
					"withdrawal": "/" + it.urlsafePathToken() + "/withdrawal",
				},
			}
		})
	}

	return func(context *gin.Context) {
		beers := getInitialBeerList()

		beers = applyQueryFilter(context, beers)

		context.JSON(200, mapHATEOAS(beers))
	}
}

func SetupApi(r *gin.Engine, temperatureProvider func() float32) {
	api := r.Group("/api")
	{
		// => first-class-functions
		api.GET("/temperature", composeGetTemperatureHandler(temperatureProvider))
		api.GET("/beers", composeGetAllBeersHandler())
	}
}

func main() {
	router := gin.Default()

	SetupApi(router, getRandomTemperature)

	router.Run()
}
