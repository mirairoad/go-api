package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	application "github.com/mirairoad/goApi/internal/app"
	"github.com/mirairoad/goApi/internal/app/config"
)

func main() {
	app := application.New(config.LoadEnv())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
