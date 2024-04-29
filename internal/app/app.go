package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mirairoad/goApi/internal/app/config"
	"github.com/mirairoad/goApi/internal/app/connectors"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	config   *config.Environment
	db       *mongo.Database
	router   http.Handler
	services *Services
}

func New(cfg *config.Environment) *App {
	app := &App{
		config: cfg,
		db:     connectors.InitDb(cfg),
	}
	app.loadServices()
	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.config.Port),
		Handler: a.router,
	}

	// redis set up
	// err := a.rdb.Ping(ctx).Err()
	// if err != nil {
	// 	return fmt.Errorf("failed to connect to redis: %w", err)
	// }

	// defer func() {
	// 	if err := a.rdb.Close(); err != nil {
	// 		fmt.Println("failed to close redis", err)
	// 	}
	// }()
	var err error = nil
	fmt.Printf("Starting server on port %s\n", a.config.Port)

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}

}
