package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var apiKey string

type WeatherInfo struct {
	Main    string
	Details struct {
		Temperature float64
	}
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey = os.Getenv("API_KEY")
}

func fetchWeatherInfo(city string) (WeatherInfo, error) {
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
		Main: data.Weather[0].Main,
	}
	weatherInfo.Details.Temperature = data.Main.Temperature

	return weatherInfo, nil
}

func main() {
	start := time.Now()

	cities := []string{"Sorocaba", "São Paulo", "Paraná", "Itu"}

	for _, city := range cities {
		weather, _ := fetchWeatherInfo(city)
		fmt.Println("Data for "+city, weather)
	}

	fmt.Println("\nTook: " + time.Since(start).String())
}
