package models

import "fmt"

type OpenWeatherModel struct {
	Country owSys       `json:"sys" :"sys" `
	City    string      `json:"name" :"city"`
	Main    owMain      `json:"main" :"main"`
	Weather []owWeather `json:"weather" :"weather"`
	Wind    owWind      `json:"speed" :"wind"`
}

type owMain struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
}

type owSys struct {
	Country string `json:"country"`
}

type owWeather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type owWind struct {
	Speed float64 `json:"speed"`
}

func (owm OpenWeatherModel) ConvertToWeatherModel() WeatherModel {
	fmt.Println(owm)
	weatherModel := WeatherModel{
		Country:     owm.Country.Country,
		City:        owm.City,
		Temp:        owm.Main.Temp,
		Pressure:    owm.Main.Pressure,
		Weather:     owm.Weather[0].Main,
		Description: owm.Weather[0].Description,
		Wind:        owm.Wind.Speed,
	}

	return weatherModel
}
