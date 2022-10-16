package forecast

import (
	"encoding/json"
	"errors"
)

var ErrInvalidJSON error = errors.New("Invalid JSON")

// New unmarshals the JSON-encoded forecast.
//
// Returns ErrInvalidJSON if invaid and/or incomplete JSON received.
func New(b []byte) (*Forecast, error) {
	var owforecast openweatherForecast
	if !json.Valid(b) {
		return nil, ErrInvalidJSON
	}

	json.Unmarshal(b, &owforecast)
	forecast := &Forecast{
		Condition:   owforecast.Condition(),
		Temperature: owforecast.Temperature(),
	}

	return forecast, nil
}

type Forecast struct {
	Condition   string
	Temperature string // use enum nstead
	Alerts      struct{}
}

type openweatherForecast struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func (o *openweatherForecast) Temperature() string {
	if o.Main.Temp > 80 {
		return "Hot"
	}

	if o.Main.Temp > 70 {
		return "Moderate"
	}

	return "Cold"
}

func (o *openweatherForecast) Condition() string {
	if len(o.Weather) == 0 {
		return "Unknown"
	}

	return o.Weather[0].Description
}
