package weather

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"weatherapp/internal/storage"
)

const apiKey = "7150d1201aa203cebf2cff874a5c2a17"

func ShowWeather(userID string) {
	users := storage.LoadUsers()
	for _, u := range users {
		if u.UserID == userID {
			location := u.Preferences.Location
			verbosity := u.Preferences.Verbosity
			forecast := u.Preferences.Forecast
			unit := u.Preferences.Unit

			if forecast == "day" {
				showCurrentWeather(location, verbosity, unit)
			} else {
				showForecastWeather(location, verbosity, forecast, unit)
			}
			return
		}
	}
}

func showCurrentWeather(location, verbosity, userUnit string) {
	url := fmt.Sprintf("http://api.weatherstack.com/current?access_key=%s&query=%s", apiKey, location)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to fetch weather data:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding weather response:", err)
		return
	}

	current, ok := result["current"].(map[string]interface{})
	if !ok {
		fmt.Println("Unexpected weather response structure")
		return
	}

	description := current["weather_descriptions"].([]interface{})[0].(string)
	temp := current["temperature"].(float64)
	feels := current["feelslike"].(float64)
	humidity := current["humidity"].(float64)
	windSpeed := current["wind_speed"].(float64)
	windDir := current["wind_dir"].(string)

	unit := "°C"
	if userUnit == "fahrenheit" {
		temp = (temp * 9 / 5) + 32
		feels = (feels * 9 / 5) + 32
		unit = "°F"
	}

	fmt.Printf("\n🌤️  Current Weather for %s\n", strings.Title(location))
	fmt.Println("----------------------------")
	fmt.Printf("Description : %s\n", description)
	fmt.Printf("Temperature : %.0f %s\n", temp, unit)

	if verbosity == "verbose" {
		fmt.Printf("Feels Like  : %.0f %s\n", feels, unit)
		fmt.Printf("Humidity    : %.0f%%\n", humidity)
		fmt.Printf("Wind        : %.0f km/h (%s)\n", windSpeed, windDir)
	}
}

func showForecastWeather(location, verbosity, forecast, unit string) {

	fmt.Printf("\n📅 Forecast for %s (%s)\n", strings.Title(location), forecast)
	fmt.Println("----------------------------")

	days := 1
	if forecast == "week" {
		days = 7
	} else if forecast == "month" {
		days = 30
	}

	unitLabel := "°C"
	if unit == "fahrenheit" {
		unitLabel = "°F"
	}

	for i := 1; i <= days; i++ {
		temp := 25 + float64(i%5)
		feels := 25 + float64(i%3)
		if unit == "fahrenheit" {
			temp = (temp * 9 / 5) + 32
			feels = (feels * 9 / 5) + 32
		}
		fmt.Printf("Day %d: %s - %.0f%s\n", i, "Partly Cloudy", temp, unitLabel)
		if verbosity == "verbose" {
			fmt.Println("  Feels like :", fmt.Sprintf("%.0f%s", feels, unitLabel))
			fmt.Println("  Humidity   : 75%")
			fmt.Println("  Wind       : 15 km/h NW")
		}
	}
}

func ShowOtherLocations(reader *bufio.Reader) {
	fmt.Print("Enter location: ")
	loc, _ := reader.ReadString('\n')
	loc = strings.TrimSpace(loc)
	showCurrentWeather(loc, "verbose", "celsius")
}
