package routes

import (
	"WeatherCaravan/controllers"
	_ "WeatherCaravan/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers.ShowStartPage)
	r.GET("/map", controllers.ShowMapPage)

	r.GET("/openWeather", controllers.GetOpenWeatherData)
	r.GET("/accuWeather", controllers.GetAccuWeatherData)
	r.GET("/weatherApi", controllers.GetWeatherApiData)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
