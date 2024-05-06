package main

import (
	"WeatherCaravan/internal/routes"
	"WeatherCaravan/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"os"
)

// @title           WeatherCaravan
// @version         1.0
// @description     Simple server for work with weather data.
func main() {

	limiter := ratelimit.New(10)

	router := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal("Ошибка загрузки переменных")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Static("/pages", "./web/pages")
	router.Static("/js", "./web/js")
	router.LoadHTMLGlob("web/pages/*")

	// Rate Limit
	router.Use(func(c *gin.Context) {
		limiter.Take()
		c.Next()
	})

	routes.SetupRoutes(router)

	addr := fmt.Sprintf(":%s", port)
	logger.Info("Запуск сервера",
		zap.String("port", port),
	)
	if err := router.Run(addr); err != nil {
		logger.Fatal("Ошибка запуска сервера",
			zap.Error(err),
		)
	}
}
