package main

import (
	"context"
	"template_backend/database"
	payment "template_backend/infrastructure/payment"
	openapi "template_backend/open-api"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Server started")
	ctx := context.Background()

	done := make(chan struct{})
	go func() {
		openapi.SetupHttp()
	}()

	database.Connect(ctx)
	payment.InitStripe()
	<-done
}
