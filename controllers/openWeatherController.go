package controllers

import (
	"WeatherCaravan/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
)

func GetOpenWeatherData(c *gin.Context) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")

	fmt.Println(latitude, longitude)

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Ошибка загрузки переменных")
		return
	}
	OWkey := os.Getenv("OWKEY")

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", latitude, longitude, OWkey)

	// Выполняем GET запрос
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении тела ответа:", err)
		return
	}
	fmt.Println(string(body))
	var owWeather models.OpenWeatherModel

	if err = json.Unmarshal(body, &owWeather); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	c.JSON(http.StatusOK, owWeather.ConvertToWeatherModel())
}
