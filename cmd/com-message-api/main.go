package main

import (
	"context"
	"fmt"
	"github.com/halink0803/zerolog-graylog-hook/graylog"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/ozonmp/com-message-api/internal/app/repo"
	"github.com/ozonmp/com-message-api/internal/app/retranslator"
	"github.com/ozonmp/com-message-api/internal/app/sender"
	"github.com/ozonmp/com-message-api/internal/config"
	"github.com/ozonmp/com-message-api/internal/database"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	log.Logger = zerolog.New(os.Stdout).With().Str("service", "com-message-api").Logger()
	hook, err := graylog.NewGraylogHook(fmt.Sprintf(cfg.Graylog.URL))
	if err != nil {
		panic(err)
	}
	//Set global logger with graylog hook
	log.Logger = log.Hook(hook)

	//migration := flag.Bool("migration", true, "Defines the migration start option")
	//flag.Parse()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	// default: zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Project.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		"localhost",
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	log.Info().Msg("DSN crom com-message-api IS:")
	log.Info().Msg(dsn)
	db, err := database.NewPostgres(dsn, cfg.Database.Driver)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
	}
	defer db.Close()

	sigs := make(chan os.Signal, 1)
	sender, error := sender.NewKafkaSender(cfg.Kafka)
	if error != nil {
		log.Fatal().Err(err).Msg("Failed init kafka sender")
	}
	rtsCfg := retranslator.Config{
		ConsumerCount:  2,
		ConsumeSize:    10,
		ConsumeTimeout: 10 * time.Second,
		ProducerCount:  28,
		WorkerCount:    2,
		ChannelSize:    512,
		Repo:           repo.NewEventRepo(db),
		Sender:         sender,
	}

	ctx, cancel := context.WithCancel(context.Background())
	retranslator := retranslator.NewRetranslator(rtsCfg)
	retranslator.Start(ctx)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	cancel()
}
