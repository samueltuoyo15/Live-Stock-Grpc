package utils

import (
	"log/slog"
	"os"
)

func InitLogger(pretty bool) *slog.Logger {
	var handler slog.Handler

	if pretty {
		handler = slog.NewTextHandler(os.Stdout, nil)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}