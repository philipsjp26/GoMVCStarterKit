package logger

import (
	"GoMVCStarterKit/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func RegisterLogger(cfg *config.EnvConfig) {
	if cfg.Application.Env == "dev" || cfg.Application.Env == "stg" {
		log, _ = zap.NewDevelopment()
	} else {
		log, _ = zap.NewProduction()
	}
	defer log.Sync()
}
func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}

// Error logs an error message
func Error(msg string, fields ...zapcore.Field) {
	log.Error(msg, fields...)
}

// Debug logs a debug message
func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}

// Warn logs a warning message
func Warn(msg string, fields ...zapcore.Field) {
	log.Warn(msg, fields...)
}

// Fatal logs a fatal message and exits the application
func Fatal(msg string, fields ...zapcore.Field) {
	log.Fatal(msg, fields...)
}
