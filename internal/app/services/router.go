package services

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mirairoad/goApi/internal/app/config"
	"github.com/mirairoad/goApi/internal/app/services/authentication"
	"github.com/mirairoad/goApi/internal/app/services/profiles"
	"github.com/mirairoad/goApi/internal/app/services/users"
)

func Load(r *chi.Mux, cfg *config.Environment) {
	//  Root Page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello, World!</h1>"))
	})

	authentication.Routes(r)
	users.Routes(r)
	profiles.Routes(r)

	// walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	// 	route = strings.Replace(route, "/*/", "/", -1)
	// 	fmt.Printf("%s %s\n", method, route)
	// 	return nil
	// }

	// if err := chi.Walk(r, walkFunc); err != nil {
	// 	fmt.Printf("Logging err: %s\n", err.Error())
	// }
}
