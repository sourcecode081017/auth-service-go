package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sourcecode081017/auth-service-go/internal/db"
)

type Handler struct {
	mongoConn *db.Mongo
}

func New(conn *db.Mongo) *Handler {
	return &Handler{
		mongoConn: conn,
	}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
