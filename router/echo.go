package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (r *Router) AnyEcho(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, struct {
		Remote      string   `json:"remote"`
		Environment []string `json:"environment"`
	}{ctx.Request.RemoteAddr, os.Environ()})
}
