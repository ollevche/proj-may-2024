package main

import (
	"time"
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
	UserID          string
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

// TODO: refactor by example
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	var u User

// 	err := json.NewDecoder(r.Body).Decode(&u)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Debug().Err(err).Msg("Failed to decode JSON")
// 		return
// 	}

// 	usersM.Lock()
// 	defer usersM.Unlock()

// 	for i := range users {
// 		if users[i].ID == u.ID {
// 			users[i] = u
// 		}
// 	}
// }
