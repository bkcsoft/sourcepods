package user

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gitpods/gitpods"
	"github.com/pressly/chi"
)

// NewUsersHandler returns a RESTful http router interacting with the Service.
func NewUsersHandler(s Service) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", list(s))
	r.Get("/:username", get(s))
	r.Put("/:username", update(s))

	return r
}

func list(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.FindAll()
		if err != nil {
			return // TODO
		}

		data, err := json.Marshal(users)
		if err != nil {
			return // TODO
		}

		w.Write(data)
	}
}

func get(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")

		user, err := s.FindByUsername(username)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return // TODO
		}

		data, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return // TODO
		}
		w.Write(data)
	}
}

func update(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")

		var user *gitpods.User
		if err := json.NewDecoder(io.LimitReader(r.Body, 5242880)).Decode(&user); err != nil {
			return // TODO
		}

		user, err := s.Update(username, user)
		if err != nil {
			return // TODO
		}

		data, err := json.Marshal(user)
		if err != nil {
			return // TODO
		}

		w.Write(data)
	}
}
