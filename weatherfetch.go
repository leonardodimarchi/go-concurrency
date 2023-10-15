package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var apiKey string

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey = os.Getenv("API_KEY")
}


func fetchWeatherInfoAsync(city string, channel chan<- WeatherInfo, waitGroup *sync.WaitGroup) (WeatherInfo, error) {
	if waitGroup != nil {
		defer waitGroup.Done()
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?units=metric&q=%s&appid=%s", city, apiKey)

	response, err := http.Get(url)

	if err != nil {
		return WeatherInfo{}, err
	}

	defer response.Body.Close()

	var data struct {
		Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
		Main struct {
			Temperature float64 `json:"temp"`
		} `json:"main"`
	}

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return WeatherInfo{}, err
	}

	weatherInfo := WeatherInfo{
		City: city,
		Main: data.Weather[0].Main,
	}
	weatherInfo.Details.Temperature = data.Main.Temperature

	if channel != nil {
		channel <- weatherInfo
	}

	return weatherInfo, nil
}