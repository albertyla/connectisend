package controller

import (
	"net/http"

	"github.com/albertyla/connectisend/internal/util"
)

func NewServiceController(logger util.Logger) *ServiceController {
	return &ServiceController{
		logger: logger,
	}
}

type ServiceController struct {
	logger util.Logger
}

// HealthHandler handles the health check endpoint.
// swagger:route GET /health HealthHandler
//
// Handles the health check request.
//
// Responses:
//
//	200: OKResponse
func (sc *ServiceController) HealthHandler(w http.ResponseWriter, r *http.Request) {
	sc.logger.InfoContext(r.Context(), "health check")
	w.WriteHeader(http.StatusOK)
}
