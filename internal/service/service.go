package service

import (
	"net/http"

	"github.com/albertyla/connectisend/internal/service/config"
	"github.com/albertyla/connectisend/internal/util"
)

func NewServer(
	logger util.Logger,
	config *config.Config,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		logger,
		config,
	)

	var handler http.Handler = mux

	return handler
}
