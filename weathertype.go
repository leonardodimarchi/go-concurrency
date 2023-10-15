package main

type WeatherInfo struct {
	City string
	Main    string
	Details struct {
		Temperature float64
	}
}