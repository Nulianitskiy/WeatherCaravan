package main

import (
	"WeatherCaravan/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/ratelimit"
	"log"
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
		fmt.Println("Ошибка загрузки переменных")
		return
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
	log.Printf("Запуск сервера на порту %s...", port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
