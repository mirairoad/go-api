package authentication

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {
	r.Route("/authentication", func(r chi.Router) {
		r.Get("/signin", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("<h1>Signin</h1>"))
		})
		r.Get("/signup", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("<h1>Signup</h1>"))
		})
		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("<h1>Logout</h1>"))
		})
	})
}
