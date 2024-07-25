package config

import (
	"context"
	"log/slog"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Host     string     `env:"HOST, default=127.0.0.1"`
	Port     string     `env:"PORT, default=8080"`
	LogLevel slog.Level `env:"LOG_LEVEL, default=INFO"`
}

func NewConfig(ctx context.Context, lookupenv func(string) (string, bool)) (*Config, error) {
	var c Config

	err := envconfig.ProcessWith(ctx, &envconfig.Config{
		Target:   &c,
		Lookuper: &wrapLookup{lookupenv},
	})

	if err != nil {
		return nil, err
	}
	return &c, nil
}

type wrapLookup struct {
	lookupenv func(string) (string, bool)
}

func (w *wrapLookup) Lookup(key string) (string, bool) {
	return w.lookupenv(key)
}
