package util

import (
	"context"
	"io"
	"log/slog"

	"github.com/albertyla/connectisend/internal/service/config"
)

type Logger interface {
	Debug(msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)
	Info(msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
	Warn(msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	Error(msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
}

func NewLogger(w io.Writer, conf *config.Config) Logger {
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: conf.LogLevel}))
}
