package main

import "net/http"

type Auth struct {
	s *Storage
}

func (a *Auth) checkAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, ok := a.s.GetUserByUsername(username)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if user.Password != password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// TODO: pass username to handler to attach trip to user
		next.ServeHTTP(w, r)
	}
}
