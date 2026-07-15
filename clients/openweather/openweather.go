package openweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenWeatherClient struct {
	apiKey string
}

func New(apiKey string) *OpenWeatherClient {
	return &OpenWeatherClient{
		apiKey: apiKey,
	}
}

func (o OpenWeatherClient) Coordinates(city string) (Coordinates, error) {
	url := "http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s"
	resp, err := http.Get(fmt.Sprintf(url, city, o.apiKey))
	if err != nil {
		return Coordinates{}, fmt.Errorf("error get coordinates: %w", err) //возврщаем ошибку, обернув в строку error get coordinates
	}

	if resp.StatusCode != 200 {
		return Coordinates{}, fmt.Errorf("error fail get coordinates: %d", resp.StatusCode)
	}

	var coordinatesResponse []CoordinatesResposne
	json.NewDecoder(resp.Body).Decode(&coordinatesResponse)
	if err != nil {
		return Coordinates{}, fmt.Errorf("error unmarshal response: %w", err)
	}

	if len(coordinatesResponse) == 0 {
		return Coordinates{}, fmt.Errorf("error empty coordinates")
	}

	return Coordinates{
		Lat: coordinatesResponse[0].Lat,
		Lon: coordinatesResponse[0].Lon,
	}, nil
}

func (o OpenWeatherClient) Weather(lat, lon float64) (Weather, error) {
	url := "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric"
	resp, err := http.Get(fmt.Sprintf(url, lat, lon, o.apiKey)) // вызываем http, который будет возвращать responce(ответ) и ошибку
	if err != nil {
		return Weather{}, fmt.Errorf("error get weather: %w", err)
	}

	if resp.StatusCode != 200 {
		return Weather{}, fmt.Errorf("error fall get weather: %d", resp.StatusCode)
	}

	var weatherResponce WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResponce)
	if err != nil {
		return Weather{}, fmt.Errorf("error unmarshal weather response: %w", err)
	}

	return Weather{
		Temp: weatherResponce.Main.Temp,
	}, nil
}
