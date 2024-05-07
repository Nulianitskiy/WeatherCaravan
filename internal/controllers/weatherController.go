package controllers

import (
	"WeatherCaravan/internal/config"
	"WeatherCaravan/internal/models"
	"WeatherCaravan/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var apiConfig config.ApiConfig

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal("Ошибка загрузки .env файла")
	}

	apiConfig.AccuWeatherAPIKey = os.Getenv("AWKEY")
	if apiConfig.AccuWeatherAPIKey == "" {
		logger.Error("Переменная AWKEY не установлена")
	}

	apiConfig.OpenWeatherAPIKey = os.Getenv("OWKEY")
	if apiConfig.OpenWeatherAPIKey == "" {
		logger.Error("Переменная OWKEY не установлена")
	}

	apiConfig.WeatherAPIAPIKey = os.Getenv("WAKEY")
	if apiConfig.WeatherAPIAPIKey == "" {
		logger.Error("Переменная WAKEY не установлена")
	}

	apiConfig.WeatherBitAPIKey = os.Getenv("WBKEY")
	if apiConfig.WeatherBitAPIKey == "" {
		logger.Error("Переменная WBKEY не установлена")
	}
}

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

	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", latitude, longitude, apiConfig.OpenWeatherAPIKey)

	resp, err := http.Get(apiUrl)
	if err != nil {
		logger.Error("Ошибка получения сообщения с API",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Ошибка при чтении тела ответа",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Println(string(body))
	var owWeather models.OpenWeatherModel

	if err = json.Unmarshal(body, &owWeather); err != nil {
		logger.Error("Ошибка при разборе JSON:",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
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

	// Place
	apiPlaceURL := "http://dataservice.accuweather.com/locations/v1/cities/geoposition/search"

	queryParams := make(map[string]string)
	queryParams["apikey"] = apiConfig.AccuWeatherAPIKey
	queryParams["q"] = fmt.Sprintf("%s,%s", latitude, longitude)

	u, _ := url.Parse(apiPlaceURL)
	q := u.Query()
	for key, value := range queryParams {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	respPlace, err := http.Get(u.String())
	if err != nil {
		logger.Error("Ошибка при выполнении запроса",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer respPlace.Body.Close()

	bodyPlace, err := ioutil.ReadAll(respPlace.Body)
	if err != nil {
		logger.Error("Ошибка при чтении ответа",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var awPlace models.AWPlace
	if err = json.Unmarshal(bodyPlace, &awPlace); err != nil {
		logger.Error("Ошибка при разборе Place JSON",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Weather
	apiWeatherURL := fmt.Sprintf("http://dataservice.accuweather.com/currentconditions/v1/%s", awPlace.PlaceKey)
	fmt.Println(awPlace.PlaceKey)
	queryParams = make(map[string]string)
	queryParams["apikey"] = apiConfig.AccuWeatherAPIKey

	u, _ = url.Parse(apiWeatherURL)
	q = u.Query()
	for key, value := range queryParams {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	respWeather, err := http.Get(u.String())
	if err != nil {
		logger.Error("Ошибка при выполнении запроса",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer respWeather.Body.Close()

	bodyWeather, err := ioutil.ReadAll(respWeather.Body)
	if err != nil {
		logger.Error("Ошибка при чтении ответа",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var awWeather []models.AWWeather
	if err = json.Unmarshal(bodyWeather, &awWeather); err != nil {
		logger.Error("Ошибка при разборе Weather JSON",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
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

	apiUrl := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?q=%s,%s&key=%s", latitude, longitude, apiConfig.WeatherAPIAPIKey)

	resp, err := http.Get(apiUrl)
	if err != nil {
		logger.Error("Ошибка загрузки API",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Ошибка при чтении тела ответа",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var waWeather models.WeatherApiModel

	if err = json.Unmarshal(body, &waWeather); err != nil {
		logger.Error("Ошибка при разборе JSON",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, waWeather.ConvertToWeatherModel())
}

// @Summary Получить информацию о погоде из WeatherBit
// @Description Получить информацию по координатам
// @Produce json
// @Param latitude query string true "Широта"
// @Param longitude query string true "Долгота"
// @Success 200 {object} models.WeatherModel
// @Router /weatherBit [get]
func GetWeatherBitData(c *gin.Context) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")

	apiUrl := fmt.Sprintf("https://api.weatherbit.io/v2.0/current?lat=%s&lon=%s&key=%s", latitude, longitude, apiConfig.WeatherBitAPIKey)

	resp, err := http.Get(apiUrl)
	if err != nil {
		logger.Error("Ошибка загрузки API",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Ошибка при чтении тела ответа",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var wbWeather models.WeatherBitModel

	if err = json.Unmarshal(body, &wbWeather); err != nil {
		logger.Error("Ошибка при разборе JSON",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	weatherModel, err := wbWeather.ConvertToWeatherModel()
	if err != nil {
		logger.Error("Ошибка при конвертации в WeatherModel",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weatherModel)
}
