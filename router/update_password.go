package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawks-atlanta/authentication-go/controller"
	"github.com/hawks-atlanta/authentication-go/models"
)

func (r *Router) UpdatePassword(ctx *gin.Context) {
	var req controller.UpdatePasswordRequest
	err := ctx.Bind(&req)
	if err != nil {
		return
	}
	session := ctx.MustGet(SessionVariale).(*models.User)
	err = r.C.UpdatePassword(session, &req)
	if err != nil {
		if errors.Is(err, controller.ErrUnauthorized) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResult)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, SucceedResult("Password updated successfully"))
}
