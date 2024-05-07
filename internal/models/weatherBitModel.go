package models

import "errors"

type WeatherBitModel struct {
	Data []WBData `json:"data"`
}

type WBData struct {
	City     string    `json:"city_name"`
	Country  string    `json:"country_code"`
	Temp     float64   `json:"temp"`
	Weather  WBWeather `json:"weather"`
	Pressure float64   `json:"pres"`
	Wind     float64   `json:"wind_speed"`
}

type WBWeather struct {
	Description string `json:"description"`
}

func (wbm WeatherBitModel) ConvertToWeatherModel() (wm WeatherModel, err error) {
	if len(wbm.Data) == 0 {
		return wm, errors.New("пустой ответ в Data")
	}

	data := wbm.Data[0]
	return WeatherModel{
		Country:     data.Country,
		City:        data.City,
		Temp:        data.Temp,
		Pressure:    data.Pressure,
		Weather:     data.Weather.Description,
		Description: data.Weather.Description,
		Wind:        data.Wind,
	}, nil
}
