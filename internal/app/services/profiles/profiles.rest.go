package profiles

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {
	r.Route("/profiles", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("<h1>all profiles</h1>"))
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			response := fmt.Sprintf("<h1>Profile ID: %s</h1>", id)
			w.Write([]byte(response))
		})
	})
}
