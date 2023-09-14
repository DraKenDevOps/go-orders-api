package main

import (
	"context"
	"fmt"

	"github.com/drakendevops/orders-api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		return fmt.Errorf("HTTP server is error: %w", err)
	}
}
