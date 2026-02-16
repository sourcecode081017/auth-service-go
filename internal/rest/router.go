package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sourcecode081017/auth-service-go/config"
)

type Router struct {
	engine *gin.Engine
	server *http.Server
}

func (h *Handler) NewRouter(cfg *config.AppConfig) *Router {
	r := gin.Default()
	r.GET("/", h.Health)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}
	return &Router{
		engine: r,
		server: server,
	}
}

func (r *Router) StartHttpServer() error {
	return r.server.ListenAndServe()
}

func (r *Router) ShutDown(ctx context.Context) error {
	return r.server.Shutdown(ctx)
}
