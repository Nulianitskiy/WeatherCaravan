package routes

import (
	_ "WeatherCaravan/docs"
	controllers2 "WeatherCaravan/internal/controllers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers2.ShowStartPage)
	r.GET("/map", controllers2.ShowMapPage)

	r.GET("/openWeather", controllers2.GetOpenWeatherData)
	r.GET("/accuWeather", controllers2.GetAccuWeatherData)
	r.GET("/weatherApi", controllers2.GetWeatherApiData)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
