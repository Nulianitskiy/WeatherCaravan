package controllers

import (
	"WeatherCaravan/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// @Summary Получить информацию о погоде из OpenWeather
// @Description Получить информацию по координатам
// @Produce json
// @Param latitude query string true "Широта"
// @Param longitude query string true "Долгота"
// @Success 200 {object} models.WeatherModel
// @Router /openWeather [get]
func GetOpenWeatherData(c *gin.Context) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Ошибка загрузки переменных")
		return
	}
	OWkey := os.Getenv("OWKEY")

	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", latitude, longitude, OWkey)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

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

// @Summary Получить информацию о погоде из AccuWeather
// @Description Получить информацию по координатам
// @Produce json
// @Param latitude query string true "Широта"
// @Param longitude query string true "Долгота"
// @Success 200 {object} models.WeatherModel
// @Router /accuWeather [get]
func GetAccuWeatherData(c *gin.Context) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Ошибка загрузки переменных")
		return
	}
	AWkey := os.Getenv("AWKEY")

	// Place
	apiPlaceURL := "http://dataservice.accuweather.com/locations/v1/cities/geoposition/search"

	queryParams := make(map[string]string)
	queryParams["apikey"] = AWkey
	queryParams["q"] = fmt.Sprintf("%s,%s", latitude, longitude)

	u, _ := url.Parse(apiPlaceURL)
	q := u.Query()
	for key, value := range queryParams {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	respPlace, err := http.Get(u.String())
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer respPlace.Body.Close()

	bodyPlace, err := ioutil.ReadAll(respPlace.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	var awPlace models.AWPlace
	if err = json.Unmarshal(bodyPlace, &awPlace); err != nil {
		fmt.Println("Ошибка при разборе Place JSON:", err)
		return
	}

	// Weather
	apiWeatherURL := fmt.Sprintf("http://dataservice.accuweather.com/currentconditions/v1/%s", awPlace.PlaceKey)
	fmt.Println(awPlace.PlaceKey)
	queryParams = make(map[string]string)
	queryParams["apikey"] = AWkey

	u, _ = url.Parse(apiWeatherURL)
	q = u.Query()
	for key, value := range queryParams {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	respWeather, err := http.Get(u.String())
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer respWeather.Body.Close()

	bodyWeather, err := ioutil.ReadAll(respWeather.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	var awWeather []models.AWWeather
	if err = json.Unmarshal(bodyWeather, &awWeather); err != nil {
		fmt.Println("Ошибка при разборе Weather JSON:", err)
		return
	}

	awModel := models.AccuWeatherModel{
		Place:   awPlace,
		Weather: awWeather[0],
	}

	c.JSON(http.StatusOK, awModel.ConvertToWeatherModel())
}

// @Summary Получить информацию о погоде из WeatherApi
// @Description Получить информацию по координатам
// @Produce json
// @Param latitude query string true "Широта"
// @Param longitude query string true "Долгота"
// @Success 200 {object} models.WeatherModel
// @Router /weatherApi [get]
func GetWeatherApiData(c *gin.Context) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Ошибка загрузки переменных")
		return
	}
	WAkey := os.Getenv("WAKEY")

	apiUrl := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?q=%s,%s&key=%s", latitude, longitude, WAkey)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении тела ответа:", err)
		return
	}

	var waWeather models.WeatherApiModel

	if err = json.Unmarshal(body, &waWeather); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}
	c.JSON(http.StatusOK, waWeather.ConvertToWeatherModel())
}
