package main

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/mirairoad/goApi/internal/app"
	"github.com/mirairoad/goApi/internal/app/config"
	"github.com/mirairoad/goApi/internal/app/connectors"
	"github.com/mirairoad/goApi/internal/app/services"
)

func main() {
	cfg := config.LoadEnv()
	connectors.DbConnect(cfg)
	r := chi.NewRouter()
	r.Use(chiMiddleware.StripSlashes)
	r.Use(chiMiddleware.Logger)
	services.Load(r, cfg)
	app.Listen(r, cfg)
}
