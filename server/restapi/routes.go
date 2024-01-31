package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *RESTAPI) routes(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	})
	router.NoMethod(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusMethodNotAllowed)
	})

	api := router.Group("/api/")
	gHash := api.Group("hash")
	{
		gHash.GET("", a.hashHandler.Get)
	}
}
