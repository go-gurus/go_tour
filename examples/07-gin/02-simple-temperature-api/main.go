package main

import (
	"github.com/gin-gonic/gin"
)

func getConstantFakeTemperature() float32 {
	return 7
}

func SetupApi(r *gin.Engine) {
	api := r.Group("/api")
	{
		// => first-class-functions
		api.GET("/temperature", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"temperature": getConstantFakeTemperature(),
			})
		})
	}
}

func main() {
	router := gin.Default()

	SetupApi(router)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
