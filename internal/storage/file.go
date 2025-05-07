package storage

import (
	"fmt"
	"os"
	"strings"
	"weatherapp/models"
)

func SaveUser(raw string) error {
	f, err := os.OpenFile("data/users.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(raw + "\n---\n")
	return err
}

func SaveAllUsers(users []models.User) error {
	f, err := os.Create("data/users.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	for _, u := range users {
		_, err = f.WriteString(SerializeUser(u) + "\n---\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func LoadUsers() []models.User {
	b, err := os.ReadFile("data/users.txt")
	if err != nil {
		return []models.User{}
	}
	blocks := strings.Split(string(b), "---\n")
	var users []models.User
	for _, block := range blocks {
		if strings.TrimSpace(block) == "" {
			continue
		}
		u := ParseUser(block)
		users = append(users, u)
	}
	return users
}

func SerializeUser(u models.User) string {
	return fmt.Sprintf(
		"userId:%s\nname:%s\npassword:%s\nlocation:%s\nunit:%s\nverbosity:%s\nforecast:%s",
		u.UserID, u.Name, u.Password,
		u.Preferences.Location, u.Preferences.Unit,
		u.Preferences.Verbosity, u.Preferences.Forecast,
	)
}

func ParseUser(raw string) models.User {
	lines := strings.Split(raw, "\n")
	u := models.User{}
	for _, l := range lines {
		parts := strings.SplitN(l, ":", 2)
		if len(parts) != 2 {
			continue
		}
		k, v := parts[0], parts[1]
		switch k {
		case "userId":
			u.UserID = v
		case "name":
			u.Name = v
		case "password":
			u.Password = v
		case "location":
			u.Preferences.Location = v
		case "unit":
			u.Preferences.Unit = v
		case "verbosity":
			u.Preferences.Verbosity = v
		case "forecast":
			u.Preferences.Forecast = v
		}
	}
	return u
}
