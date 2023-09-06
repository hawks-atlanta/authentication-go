package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawks-atlanta/authentication-go/controller"
	"github.com/hawks-atlanta/authentication-go/models"
)

type Token struct {
	JWT string `json:"jwt"`
}

// POST /login
func (r *Router) Login(ctx *gin.Context) {
	var creds models.User
	err := ctx.Bind(&creds)
	if err != nil {
		return
	}
	user, err := r.C.Login(&creds)
	if err != nil {
		if errors.Is(err, controller.ErrUnauthorized) {
			ctx.JSON(http.StatusUnauthorized, UnauthorizedResult)
		} else {
			ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
		}
		return
	}
	ctx.JSON(http.StatusCreated, Token{JWT: r.C.JWT.New(user.Claims())})
}
