package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", CreateUser)
	mux.HandleFunc("PUT /users", UpdateUser)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}
