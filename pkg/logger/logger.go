package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var baseLogger = newLogger()

func newLogger() *zap.SugaredLogger {
	atomicLevel := zap.NewAtomicLevel()

	atomicLevel.SetLevel(zap.InfoLevel)
	if debugStr := os.Getenv("DEBUG"); debugStr == "true" {
		atomicLevel.SetLevel(zap.DebugLevel)
	}

	zapCfg := zap.NewProductionConfig()
	zapCfg.Level = atomicLevel
	zapCfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	z, err := zapCfg.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	return z.Sugar()
}

// Init should be called in the main entrypoint to initialize the logger with common logger fields (version, commit, buildDate).
func Init(args ...any) *zap.SugaredLogger {
	baseLogger = baseLogger.With(args...)
	return baseLogger
}

// Logger returns a copy of the base logger with provided logger fields.
func Logger(args ...any) *zap.SugaredLogger {
	return baseLogger.With(args...)
}
