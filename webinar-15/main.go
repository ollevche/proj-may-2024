package main

import (
	"net/http"

	"w15/internal/user"

	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	userStorage := user.NewInMemStorage()

	userService := user.NewService(userStorage)

	userHandler := user.NewHandler(userService)

	mux.HandleFunc("POST /users", userHandler.Create)
	// mux.HandleFunc("PUT /users", UpdateUser)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}
