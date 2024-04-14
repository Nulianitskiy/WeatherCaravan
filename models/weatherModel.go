package models

type WeatherModel struct {
	Country     string  `json:"country" :"country" `
	City        string  `json:"name" :"city"`
	Temp        float32 `json:"temp" :"temp"`
	Pressure    float32 `json:"pressure"`
	Weather     string  `json:"main" :"weather"`
	Description string  `json:"description" :"description"`
	Wind        float32 `json:"speed" :"wind"`
}
