package main

import (
	"fmt"
	"generic-backend-service/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Hello World")
	cfg := config.GetDefault()
	ConfigureLogger(cfg.LoggerInfo)

}

func ConfigureLogger(info *config.LoggerInfo) {
	zerolog.SetGlobalLevel(info.LogLevel)
	log.Log().Msg("finished configuring logger")
}
