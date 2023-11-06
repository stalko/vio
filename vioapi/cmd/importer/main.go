package main

import (
	"context"
	"log"

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

	CountGoRoutine  int `env:"COUNT_GO_ROUTINE"`
	CountBulkInsert int `env:"COUNT_BULK_INSERT"`

	DatabaseDSN string `env:"DATABASE_DSN"`
}

func main() {
	cfg, err := config.Load[Config]()
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	logger := logging.CreateLogger(cfg.Environment, cfg.SentryDSN, nil)

	logger.Info("importer service - started")

	ctx := context.Background()

	s, err := db.NewDB(ctx, cfg.DatabaseDSN, logger)
	if err != nil {
		log.Fatalf("Failed to create db connection: %s", err.Error())
	}

	imp := importer.NewCSVImporter(s, cfg.CountBulkInsert, logger, ctx)

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
