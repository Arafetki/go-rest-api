package main

import (
	"log/slog"
	"os"
	"runtime/debug"
	"strings"

	"github.com/Arafetki/my-portfolio-api/internal/env"
	"github.com/lmittmann/tint"
)

type config struct {
	httpPort int
	env      string
}

type application struct {
	cfg    config
	logger *slog.Logger
}

func main() {
	var cfg config
	var logLevel slog.Level
	cfg.httpPort = env.GetInt("APP_PORT", 8080)
	cfg.env = env.GetString("APP_ENV", "development")
	switch strings.ToLower(cfg.env) {
	case "staging":
		logLevel = slog.LevelInfo
	case "production":
		logLevel = slog.LevelWarn
	default:
		logLevel = slog.LevelDebug
	}

	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: logLevel}))

	app := &application{
		cfg:    cfg,
		logger: logger,
	}

	err := app.serveHTTP()
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}
