package main

import (
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type envelope map[string]any

func getIDParam(r *http.Request) (int, error) {
	idParam := chi.URLParamFromCtx(r.Context(), "id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return 0, errors.New("bad id param")
	}
	return id, nil
}

func parseLogLevel(logLevelStr string) slog.Level {
	var logLevel slog.Level
	switch strings.ToLower(logLevelStr) {
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelDebug
	}

	return logLevel
}

func readString(qs url.Values, key string, defaultValue string) string {

	value := qs.Get(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func readInt(qs url.Values, key string, defaultValue int) int {

	stringValue := qs.Get(key)

	if stringValue == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(stringValue)
	if err != nil {
		return defaultValue
	}

	return intValue
}

func readCsv(qs url.Values, key string, defaultValue []string) []string {

	csv := qs.Get(key)

	if csv == "" {
		return defaultValue
	}

	return strings.Split(csv, ",")
}
