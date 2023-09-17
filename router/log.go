package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /log
func (r *Router) Log(ctx *gin.Context) {

	logs, err := r.C.GetLogs()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}
	if len(logs) == 0 {
		ctx.JSON(http.StatusNotFound, InternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, logs)
}
