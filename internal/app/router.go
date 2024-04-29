package application

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chiMid "github.com/go-chi/chi/v5/middleware"
	"github.com/mirairoad/goApi/internal/app/services/authentication"
	"github.com/mirairoad/goApi/internal/app/services/users"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()
	router.Use(chiMid.Logger)

	//  REST
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello, World!</h1>"))
	})

	authentication.Routes(router)
	router.Route("/users", a.userRoutes)

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>404, not found!</h1>"))
	})

	a.router = router
}

func (a *App) userRoutes(router chi.Router) {
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var data *users.User

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest) // Set status code to 400
			json.NewEncoder(w).Encode(bson.M{"message": "Bad Request"})
			fmt.Println(err.Error())
			return
		}

		if data.Email == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest) // Set status code to 400
			json.NewEncoder(w).Encode(bson.M{"message": "Email is required"})
			fmt.Println(bson.M{"error": "Email is required"})
			return
		}

		if data.Password == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest) // Set status code to 400
			json.NewEncoder(w).Encode(bson.M{"message": "Password is required"})
			fmt.Println(bson.M{"error": "Password is required"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data.Password = string(hashedPassword)

		user, err := a.services.users.Create(data, ctx)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})
}
