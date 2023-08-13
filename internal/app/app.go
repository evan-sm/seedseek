package app

import (
	"context"
	"os"
	"os/signal"

	"seedseek/internal/config"
	"seedseek/internal/indexer"
	"seedseek/internal/infrastructure/bot"
	"seedseek/pkg/logger"

	"github.com/go-redis/redis"
	"golang.org/x/exp/slog"
)

func Run(cfg *config.Config) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	log := logger.New()

	log.DebugContext(ctx, "Hello world!", slog.Int("test", 123))

	// Databases and caches
	r := redis.NewClient(&redis.Options{
		Addr: cfg.App.RedisAddr,
	})

	// Repositories

	// Adapters

	// Indexers
	indexers := indexer.New(log, cfg, r)

	// Services

	// Daemon
	// go daemon.New(log).Run(ctx)

	// HTTP Server

	// Bot

	botSvc, err := bot.New(ctx, cfg, log, indexers)
	if err != nil {
		log.ErrorContext(ctx, "app - Run - bot.New: %w", err)
		os.Exit(1)
	}

	botSvc.Run(ctx)

	<-ctx.Done()
	log.InfoContext(ctx, "Gracefully shutting down...")
}
