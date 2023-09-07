package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawks-atlanta/authentication-go/controller"
)

func (r *Router) Authorize(ctx *gin.Context) {
	header := ctx.GetHeader(AuthorizationHeader)
	if len(header) < 7 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResult)
		return
	}
	user, err := r.C.Authorize(header[7:])
	if err != nil {
		if errors.Is(err, controller.ErrUnauthorized) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResult)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}
	ctx.Set(SessionVariale, &user)
	ctx.Next()
}
