package main

import (
	"WeatherCaravan/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
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

	routes.SetupRoutes(router)

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Запуск сервера на порту %s...", port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
