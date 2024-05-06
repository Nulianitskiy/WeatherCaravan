package models

type WeatherApiModel struct {
	Location waLocation `json:"location"`
	Current  waCurrent  `json:"current"`
}

type waLocation struct {
	Country string `json:"country"`
	City    string `json:"region"`
}

type waCurrent struct {
	Temp      float64     `json:"temp_c"`
	Condition waCondition `json:"condition"`
	Pressure  float64     `json:"pressure_mb"`
	Speed     float64     `json:"wind_kph"`
}

type waCondition struct {
	Description string `json:"text"`
}

func (wam WeatherApiModel) ConvertToWeatherModel() WeatherModel {
	weatherModel := WeatherModel{
		Country:     wam.Location.Country,
		City:        wam.Location.City,
		Temp:        wam.Current.Temp,
		Pressure:    wam.Current.Pressure,
		Weather:     wam.Current.Condition.Description,
		Description: "",
		Wind:        wam.Current.Speed,
	}
	return weatherModel
}
