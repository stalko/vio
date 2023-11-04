package main

import (
	"context"
	"log"
	"time"

	"github.com/stalko/vioapi/pkg/config"
	"github.com/stalko/vioapi/pkg/logging"
	"github.com/stalko/viodata/db"
	"github.com/stalko/viodata/importer"
	"go.uber.org/zap"
)

type Config struct {
	Environment string `env:"ENV" envDefault:"development" validate:"required"`
	SentryDSN   string `env:"SENTRY_DSN"`

	ImportFile string `env:"IMPORT_FILE"`

	WorkerTimeoutSec int64 `env:"WORKER_TIMEOUT_SEC"`
	CountGoRoutine   int   `env:"COUNT_GO_ROUTINE"`

	DatabaseDSN          string `env:"DATABASE_DSN"`
	DBConnMaxIdleTimeSec int    `env:"DB_CONN_MAX_IDLE_TIME_SEC"`
	DBBackoffDurationSec int    `env:"DB_BACKOFF_DURATION_SEC"`
	DBMaxOpenConns       int    `env:"DB_MAX_OPEN_CONNS"`
	DBBackoffRetryCount  uint64 `env:"DB_BACKOFF_RETRY_COUNT"`
}

func main() {
	cfg, err := config.Load[Config]()
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	logger := logging.CreateLogger(cfg.Environment, cfg.SentryDSN, nil)

	logger.Info("importer service - started")

	ctx := context.Background()

	s, err := db.NewDB(ctx, cfg.DatabaseDSN, logger, db.DBConfig{
		ConnMaxIdleTime:   time.Duration(cfg.DBConnMaxIdleTimeSec) * time.Second,
		MaxOpenConns:      cfg.DBMaxOpenConns,
		BackoffRetryCount: cfg.DBBackoffRetryCount,
		BackoffDuration:   time.Duration(cfg.DBBackoffDurationSec) * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to create db connection: %s", err.Error())
	}

	imp := importer.NewCSVImporter(s, time.Duration(cfg.WorkerTimeoutSec)*time.Second)

	result, err := imp.Import(cfg.ImportFile, cfg.CountGoRoutine)
	if err != nil {
		log.Fatalf("Failed to import data: %s", err.Error())
	}

	logger.Info("Import - finished",
		zap.Duration("importing_duration", result.Duration),
		zap.Int("accepted_entries", result.AcceptedEntries),
		zap.Int("discarded_entries", result.DiscardedEntries),
	)

	logger.Info("program exited")
}
