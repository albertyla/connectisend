package service

import (
	"net/http"

	"github.com/albertyla/connectisend/internal/service/config"
	"github.com/albertyla/connectisend/internal/service/controller"
	"github.com/albertyla/connectisend/internal/util"
)

func addRoutes(
	mux *http.ServeMux,
	logger util.Logger,
	_ *config.Config,
) {
	sc := controller.NewServiceController(logger)

	mux.HandleFunc("/health", sc.HealthHandler)
}
