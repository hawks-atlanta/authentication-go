package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawks-atlanta/authentication-go/internal/utils/ipaddr"
	"github.com/hawks-atlanta/authentication-go/models"
)

func (r *Router) Challenge(ctx *gin.Context) {
	session := ctx.MustGet(SessionVariale).(*models.User)
	log := models.Log{User: session, UserUUID: session.UUID, Action: "User JWT renewal", IpAddress: ipaddr.GetIpAddr(ctx)}
	err := r.C.CreateLog(&log)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, Token{JWT: r.C.JWT.New(session.Claims())})
}
