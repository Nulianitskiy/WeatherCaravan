package models

type WeatherModel struct {
	Country     string  `json:"country" :"country" `
	City        string  `json:"name" :"city"`
	Temp        float64 `json:"temp" :"temp"`
	Pressure    float64 `json:"pressure"`
	Weather     string  `json:"main" :"weather"`
	Description string  `json:"description" :"description"`
	Wind        float64 `json:"speed" :"wind"`
}
