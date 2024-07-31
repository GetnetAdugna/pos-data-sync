package services

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	zerolog "github.com/rs/zerolog/log"
)

func StartScheduler() {
	s := gocron.NewScheduler(time.UTC)

	_, err := s.Every(1).Minute().Do(syncJob)
	if err != nil {
		log.Fatalf("Error scheduling job: %v", err)
	}

	s.StartAsync()
}

func syncJob() {
	zerolog.Debug().Msg("Running sync job")
	// Add logic to fetch transactions and publish to MQTT broker here
	FetchAndPublishTransactions()
}
