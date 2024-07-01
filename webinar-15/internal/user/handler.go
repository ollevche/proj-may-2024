package user

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type CreateUserReqBody struct {
	Email    string
	Password string
}

type service interface {
	SignUp(email, password string)
}

type Handler struct {
	s service
}

func NewHandler(s service) Handler {
	return Handler{s: s}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateUserReqBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Debug().Err(err).Msg("Failed to decode JSON")
		return
	}

	h.s.SignUp(reqBody.Email, reqBody.Password)
}
