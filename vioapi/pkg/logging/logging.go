package logging

import (
	"log"

	"github.com/TheZeroSlave/zapsentry"
	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CreateLogger using sentry depending on the environment
// It is possible to overwrite the sentry configuration by passing a variable of type zapsentry.Configuration to this funtion. If you wish to use the default configuration, pass nil
func CreateLogger(env, sentryDSN string, sentryCfg *zapsentry.Configuration) *zap.Logger {
	switch env {
	case "production":
		if sentryCfg == nil {
			sentryCfg = &zapsentry.Configuration{
				Level: zapcore.ErrorLevel,
				Tags: map[string]string{
					"env": env,
				},
			}
		}

		sentryCore, err := zapsentry.NewCore(*sentryCfg, zapsentry.NewSentryClientFromDSN(sentryDSN))
		if err != nil {
			log.Fatalf("Failed to create sentry core: %s", err.Error())
		}

		logger, err := zapdriver.NewProduction()
		if err != nil {
			log.Fatalf("Failed to create logger: %s", err.Error())
		}

		return zapsentry.AttachCoreToLogger(sentryCore, logger)

	case "staging":
		logger, err := zapdriver.NewDevelopment()
		if err != nil {
			log.Fatalf("Failed to create logger: %s", err.Error())
		}

		return logger
	}

	// Development
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Failed to create logger: %s", err.Error())
	}

	return logger
}
