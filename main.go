package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchAllSync() {
	cities := []string{"Sorocaba", "São Paulo", "Paraná", "Itu"}

	for _, city := range cities {
		weather, _ := fetchWeatherInfoAsync(city, nil, nil)
		fmt.Println("Data for "+city, weather)
	}
}

func fetchAllAsync() {
	cities := []string{"Sorocaba", "São Paulo", "Paraná", "Itu"}

	channel := make(chan WeatherInfo)
	var waitGroup sync.WaitGroup

	for _, city := range cities {
		waitGroup.Add(1)
		go fetchWeatherInfoAsync(city, channel, &waitGroup)
	}

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	for result := range channel {
		fmt.Println("Data for "+result.City, result)
	}
}

func main() {
	fmt.Println("------Sync------")
	startSync := time.Now()

	fetchAllSync()

	fmt.Println("\nTook: " + time.Since(startSync).String())

	fmt.Println("\n------Async------")

	startAsync := time.Now()

	fetchAllAsync()

	fmt.Println("\nTook: " + time.Since(startAsync).String())
}
