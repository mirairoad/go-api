package users

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	application "github.com/mirairoad/goApi/internal/app"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"golang.org/x/crypto/bcrypt"
// )

// func Routes(a *application.App) {
// 	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()

// 		var data *User

// 		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusBadRequest) // Set status code to 400
// 			json.NewEncoder(w).Encode(bson.M{"message": "Bad Request"})
// 			fmt.Println(err.Error())
// 			return
// 		}

// 		if data.Email == "" {
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusBadRequest) // Set status code to 400
// 			json.NewEncoder(w).Encode(bson.M{"message": "Email is required"})
// 			fmt.Println(bson.M{"error": "Email is required"})
// 			return
// 		}

// 		if data.Password == "" {
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusBadRequest) // Set status code to 400
// 			json.NewEncoder(w).Encode(bson.M{"message": "Password is required"})
// 			fmt.Println(bson.M{"error": "Password is required"})
// 			return
// 		}

// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

// 		if err != nil {
// 			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 			return
// 		}

// 		data.Password = string(hashedPassword)

// 		user, err := Create(data, ctx)

// 		if err != nil {
// 			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(user)
// 	})

// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	// var results []bson.M
// 	users, err := Find(ctx)

// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(users)
// })

// r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	id := chi.URLParam(r, "id")
// 	user, err := Get(id, ctx)

// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// })

// r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	var data bson.M

// 	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusBadRequest) // Set status code to 400
// 		json.NewEncoder(w).Encode(bson.M{"error": err.Error()})
// 		return
// 	}

// 	id := chi.URLParam(r, "id")
// 	user, err := Patch(id, data, ctx)

// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// })

// r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	id := chi.URLParam(r, "id")

// 	user, err := Delete(id, ctx)

// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// })
// }
