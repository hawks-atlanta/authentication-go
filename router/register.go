package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hawks-atlanta/authentication-go/controller"
	"github.com/hawks-atlanta/authentication-go/internal/utils/ipaddr"
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
		if strings.Contains(err.Error(), controller.ErrDuplicatedUSer.Error()) {
			ctx.AbortWithStatusJSON(http.StatusConflict, DuplicatedUserResult)
			return
		}
		ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}
	log := models.Log{User: &user, UserUUID: user.UUID, Action: "User registration", IpAddress: ipaddr.GetIpAddr(ctx)}
	err = r.C.CreateLog(&log)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}
	ctx.JSON(http.StatusCreated, Token{JWT: r.C.JWT.New(user.Claims())})
}
