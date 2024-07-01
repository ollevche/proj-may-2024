package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type Address struct {
	City     string
	Street   string
	Building string
	Flat     string
	Comment  string
}

type Item struct {
	Product
	Amount int
}

type Order struct {
	ID              string
	Status          string
	Items           []Item
	DeliveryAddress Address
	CreatedAt       time.Time
	TotalAmount     int
}

type Product struct {
	ID          string
	Name        string
	Description string
	Price       int
}

type User struct {
	ID             string
	Email          string
	Password       string
	SignedUpAt     time.Time
	Orders         []Order
	Role           string
	ViewedProducts []Product
}

var (
	usersM sync.Mutex
	users  []User
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Debug().Err(err).Msg("Failed to decode JSON")
		return
	}

	usersM.Lock()
	defer usersM.Unlock()
	users = append(users, u)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Debug().Err(err).Msg("Failed to decode JSON")
		return
	}

	usersM.Lock()
	defer usersM.Unlock()

	for i := range users {
		if users[i].ID == u.ID {
			users[i] = u
		}
	}
}
