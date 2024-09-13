package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const APIKEY string = "bf62e4c080ca4d2450eccd00b94db999"

type Weather struct {
	Name string `json:"name"`
	Main struct {
		Temp     float32 `json: "temp"`
		Pressure int     `json: "pressure"`
		Humidity int     `json: "humidity"`
	} `json "main"`
	Weather []struct {
		Description string `json:"description"`
	} `json: "weather"`
}

func getWeather(city string) (*Weather, error) {
	// apiUrl :=
	reqUrl := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, APIKEY)
	fmt.Println("reqUrl: ", reqUrl)
	resp, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weatherResp Weather
	err1 := json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err1 != nil {
		return nil, err
	}

	return &weatherResp, nil
}

func main() {

	envHomePath := strings.Split(os.Getenv("HOMEPATH"), "\\")
	envName := envHomePath[len(envHomePath)-1]
	fmt.Print("Hello, ", envName, " \nenter your city: ")
	var city string
	fmt.Scan(&city)
	currentWeather, err := getWeather(city)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Current weather in %s:\n", currentWeather.Name)
	fmt.Printf("Temperature: %.2f°C\n", currentWeather.Main.Temp)
	fmt.Printf("Pressure: %d hPa\n", currentWeather.Main.Pressure)
	fmt.Printf("Humidity: %d%%\n", currentWeather.Main.Humidity)
	fmt.Printf("Description: %s\n", currentWeather.Weather[0].Description)

	var a int
	fmt.Scan(&a)
}
