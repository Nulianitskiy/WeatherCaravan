package models

type AccuWeatherModel struct {
	Place   AWPlace
	Weather AWWeather
}

type AWPlace struct {
	PlaceKey string    `json:"Key"`
	Country  awCountry `json:"Country"`
	City     awCity    `json:"ParentCity"`
}

type awCountry struct {
	Country string `json:"ID"`
}

type awCity struct {
	City string `json:"EnglishName"`
}

type AWWeather struct {
	Weather     string `json:"PrecipitationType"`
	Description string `json:"WeatherText"`
	Temp        awTemp `json:"Temperature"`
}
type awTemp struct {
	Temp awMetric `json:"Metric"`
}
type awMetric struct {
	Value float32 `json:"Value"`
}

func (awm AccuWeatherModel) ConvertToWeatherModel() WeatherModel {
	weatherModel := WeatherModel{
		Country:     awm.Place.Country.Country,
		City:        awm.Place.City.City,
		Temp:        awm.Weather.Temp.Temp.Value,
		Pressure:    -1,
		Weather:     awm.Weather.Weather,
		Description: awm.Weather.Description,
		Wind:        -1,
	}

	return weatherModel
}
