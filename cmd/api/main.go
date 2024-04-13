package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"github.com/Arafetki/my-portfolio-api/internal/env"
	"github.com/Arafetki/my-portfolio-api/internal/secrets"
	"github.com/lmittmann/tint"
)

type config struct {
	httpPort int
	env      string
	db       struct {
		dsn string
	}
}

type application struct {
	cfg     config
	logger  *slog.Logger
	secrets *secrets.Secrets
	wg      sync.WaitGroup
}

const version = "1.0.0"

func main() {

	var cfg config
	cfg.httpPort = env.GetInt("APP_PORT", 8080)
	cfg.env = env.GetString("APP_ENV", "development")

	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))

	v, err := secrets.NewVault("secret")
	if err != nil {
		logger.Error(fmt.Sprintf("vault: %s", err.Error()))
		os.Exit(1)
	}
	logger.Info("vault: access granted")

	cfg.db.dsn = v.GetSecret("database")["dsn"].(string)

	app := &application{
		cfg:     cfg,
		logger:  logger,
		secrets: secrets.New(v),
	}

	err = app.serveHTTP()
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}
