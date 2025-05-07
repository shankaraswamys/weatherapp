package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"weatherapp/internal/auth"
	"weatherapp/internal/user"
	"weatherapp/internal/weather"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Weather CLI App ===")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			auth.Register(reader)
		case "2":
			userID := auth.Login(reader)
			if userID != "" {
				user.EnsurePreferences(reader, userID)
				dashboard(reader, userID)
			}
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func dashboard(reader *bufio.Reader, userID string) {
	for {
		fmt.Println("\n=== Dashboard ===")
		fmt.Println("1. View My Weather")
		fmt.Println("2. Change Preferences")
		fmt.Println("3. View Other Locations")
		fmt.Println("4. List Users")
		fmt.Println("5. Logout")
		fmt.Print("Enter choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			weather.ShowWeather(userID)
		case "2":
			user.ChangePreferences(reader, userID)
		case "3":
			weather.ShowOtherLocations(reader)
		case "4":
			user.ListUsers()
		case "5":
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
