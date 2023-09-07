package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawks-atlanta/authentication-go/models"
)

func (r *Router) Challenge(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Token{JWT: r.C.JWT.New(ctx.MustGet(SessionVariale).(*models.User).Claims())})
	return
}
