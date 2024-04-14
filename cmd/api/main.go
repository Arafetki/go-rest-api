package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	database "github.com/Arafetki/my-portfolio-api/internal/db"
	"github.com/Arafetki/my-portfolio-api/internal/env"
	"github.com/Arafetki/my-portfolio-api/internal/secrets"
	"github.com/Arafetki/my-portfolio-api/internal/vault"
	"github.com/lmittmann/tint"
)

type config struct {
	httpPort int
	env      string
	db       struct {
		dsn         string
		automigrate bool
	}
}

type application struct {
	cfg         config
	logger      *slog.Logger
	secretStore *secrets.Store
	wg          sync.WaitGroup
}

const version = "1.0.0"

func main() {

	var cfg config
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))

	// Establish Vault Connection
	v, err := vault.NewVault("secret")
	if err != nil {
		logger.Error(fmt.Sprintf("vault: %s", err.Error()))
		os.Exit(1)
	}
	// Initialize Secret Store
	secretStore := secrets.NewStore(v)

	cfg.db.dsn, err = secretStore.Provider.ReadString("database", "dsn")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	cfg.httpPort = env.GetInt("APP_PORT", 8080)
	cfg.env = env.GetString("APP_ENV", "development")
	cfg.db.automigrate = env.GetBool("DB_AUTOMIGRATE", true)

	db, err := database.New(cfg.db.dsn, cfg.db.automigrate)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("db connection has been established sucessfully!")

	app := &application{
		cfg:         cfg,
		logger:      logger,
		secretStore: secretStore,
	}

	err = app.serveHTTP()
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}
