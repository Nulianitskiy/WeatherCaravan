package routes

import (
	"WeatherCaravan/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers.ShowStartPage)
	r.GET("/openWeather", controllers.GetOpenWeatherData)
}
