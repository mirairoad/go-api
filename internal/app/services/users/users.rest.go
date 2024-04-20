package users

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("<h1>all users</h1>"))
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			response := fmt.Sprintf("<h1>User ID: %s</h1>", id)
			w.Write([]byte(response))
		})
	})
}
