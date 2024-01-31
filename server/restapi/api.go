package restapi

import (
	"context"
	"net/http"

	"github.com/dolefir/refresh-hash/config"
	"github.com/dolefir/refresh-hash/logger"
	hashs "github.com/dolefir/refresh-hash/server/restapi/handlers"
	"github.com/gin-gonic/gin"
)

// RESTAPI encapsulates necessary dependencies
// for running server.
type RESTAPI struct {
	hashHandler *hashs.Handler
	cfg         config.APIServer
	log         logger.Logger
	srv         http.Server
}

// NewAPI returns a new REST API with dependencies.
func NewAPI(hashHandler *hashs.Handler, cfg config.APIServer, log logger.Logger) *RESTAPI {
	api := &RESTAPI{
		hashHandler: hashHandler,
		cfg:         cfg,
		log:         log,
	}
	router := gin.Default()
	api.routes(router)

	api.srv = http.Server{
		Addr:    api.cfg.HTTP.ListenAddr,
		Handler: router,
	}

	return api
}

// ListenAndServe starts an API server.
func (a *RESTAPI) ListenAndServe(ctx context.Context) error {
	return a.srv.ListenAndServe()
}

func (a *RESTAPI) Shutdown(ctx context.Context) error {
	return a.srv.Shutdown(ctx)
}
