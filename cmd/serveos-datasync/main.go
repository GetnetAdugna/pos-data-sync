package main

import (
	"flag"
	"fmt"

	"serveos-datasync/code/api"
	"serveos-datasync/code/business"
	"serveos-datasync/code/db"
	"serveos-datasync/code/services"
	"serveos-datasync/code/services/mqtt"
	"serveos-datasync/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug logging")
	flag.Parse()

	services.InitLogging(*debug)

	cfg := config.InitConfig()

	db.InitDB(cfg)

	license, err := business.VerifyLicense()
	if err != nil {
		log.Fatal().Err(err).Msg("License verification failed")
	}

	mqtt.InitMQTT(cfg)

	services.StartScheduler()

	fmt.Printf("Loaded configuration: %+v\n", cfg)

	log.Info().Str("MQTT Host", cfg.ServerMQTTHost).Str("Database Name", cfg.DatabaseName).Interface("License", license).Msg("Configuration values")

	router := gin.Default()

	api.InitStatusRouter(router)

	router.Run(":8080")
	select {}
}
