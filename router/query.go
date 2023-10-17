package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawks-atlanta/authentication-go/controller"
	"github.com/hawks-atlanta/authentication-go/internal/utils/ipaddr"
	"github.com/hawks-atlanta/authentication-go/models"
)

func (r *Router) UserByUsername(ctx *gin.Context) {
	var req = controller.UserRequest{
		Username: ctx.Param(UsernameParam),
	}
	user, err := r.C.UserByUsername(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	log := models.Log{User: &user, UserUUID: user.UUID, Action: "Got user by username", IpAddress: ipaddr.GetIpAddr(ctx)}
	err = r.C.CreateLog(&log)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	res := models.User{Model: user.Model, Username: user.Username}
	ctx.JSON(http.StatusOK, res)
}

func (r *Router) UsernameByUUID(ctx *gin.Context) {
	var req = controller.UserRequest{
		UUID: ctx.Param(UUIDParam),
	}
	user, err := r.C.UsernameByUUID(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	log := models.Log{User: &user, UserUUID: user.UUID, Action: "Got username by uuid", IpAddress: ipaddr.GetIpAddr(ctx)}
	err = r.C.CreateLog(&log)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	res := models.User{Model: user.Model, Username: user.Username}
	ctx.JSON(http.StatusOK, res)
}
