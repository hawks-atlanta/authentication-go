package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawks-atlanta/authentication-go/models"
)

func (r *Router) Register(ctx *gin.Context) {
	var user models.User
	err := ctx.Bind(&user)
	if err != nil {
		return
	}
	err = r.C.Register(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}
	ctx.JSON(http.StatusCreated, Token{JWT: r.C.JWT.New(user.Claims())})
}
