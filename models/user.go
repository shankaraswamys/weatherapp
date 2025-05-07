package models

type Preferences struct {
	Location  string
	Unit      string
	Verbosity string
	Forecast  string
}

type User struct {
	UserID      string
	Name        string
	Password    string
	Preferences Preferences
}
