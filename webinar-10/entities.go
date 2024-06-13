package main

type Trip struct {
	ID                 int
	Title              string
	Description        string
	Source             string
	Destination        string
	DestinationWeather string
	// Username    string
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
