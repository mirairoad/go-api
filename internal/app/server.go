package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mirairoad/goApi/internal/app/config"
)

func Listen(r *chi.Mux, cfg *config.Environment) {
	port := fmt.Sprintf(":%s", cfg.Port)
	url := fmt.Sprintf("%s://%s:%s", cfg.Protocol, cfg.Host, cfg.Port)

	fmt.Printf("Server is running on %s\n", url)

	http.ListenAndServe(port, r)
}
