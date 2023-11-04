package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/stalko/vioapi/pkg/config"
	"github.com/stalko/vioapi/pkg/logging"
	"github.com/stalko/vioapi/pkg/server"
	"github.com/stalko/viodata"
	"github.com/stalko/viodata/db"
	"go.uber.org/zap"
)

type Config struct {
	Environment string `env:"ENV" envDefault:"development" validate:"required"`
	HTTPPort    string `env:"HTTP_PORT" envDefault:"8080"`
	SentryDSN   string `env:"SENTRY_DSN"`

	ShutdownGracePeriodSeconds int `env:"SHUTDOWN_GRACE_PERIOD_SECONDS" envDefault:"5"`

	DatabaseDSN string `env:"DATABASE_DSN"`
}

func main() {
	cfg, err := config.Load[Config]()
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	logger := logging.CreateLogger(cfg.Environment, cfg.SentryDSN, nil)
	logger.Info("api service - started")

	// setup graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s, err := db.NewDB(ctx, cfg.DatabaseDSN, logger, db.DBConfig{
		ConnMaxIdleTime:   30 * time.Second,
		MaxOpenConns:      10,
		BackoffRetryCount: 3,
		BackoffDuration:   3 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to create db connection: %s", err.Error())
	}

	vd := viodata.NewVioData(s, logger)

	srv := server.NewHTTPServer(cfg.HTTPPort, logger, vd)

	srv.Run()

	<-ctx.Done()
	err = ctx.Err()
	logger.Info("os signal received, terminating gracefully. send signal again to force shutdown", zap.Error(err))

	srv.Shutdown(ctx)
	stop()

	timeout := cfg.ShutdownGracePeriodSeconds
	if timeout > 0 {
		logger.Info("received shutdown signal", zap.Int("timeout", timeout))
		time.Sleep(time.Second * time.Duration(timeout))
	}

	logger.Info("program exited")
}
