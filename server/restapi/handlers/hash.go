package handlers

import (
	"net/http"

	"github.com/dolefir/refresh-hash/models"
	"github.com/dolefir/refresh-hash/services"
	"github.com/gin-gonic/gin"
)

// Handler holds all actions for hash.
type Handler struct {
	hashSrv services.Hash
}

// NewHandler return a new handler.
func NewHandler(hash services.Hash) *Handler {
	return &Handler{
		hashSrv: hash,
	}
}

// Get - handler GET for /api/hash endpoint.
func (h Handler) Get(ctx *gin.Context) {
	hash, err := h.hashSrv.Get(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, models.Hash{ID: hash.ID, Datatime: hash.Datatime})
}
