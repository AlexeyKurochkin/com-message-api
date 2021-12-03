package main

import (
	"github.com/ozonmp/com-message-api/internal/config"
	"github.com/rs/zerolog/log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	routerPkg "github.com/ozonmp/com-message-api/internal/app/omp-bot/router"
)

func main() {
	_ = godotenv.Load()
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Fatal().Msg("environment variable TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal().Err(err)
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal().Err(err)
	}

	routerHandler := routerPkg.NewRouter(bot, &cfg)

	for update := range updates {
		routerHandler.HandleUpdate(update)
	}
}
