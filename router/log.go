package router

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hawks-atlanta/authentication-go/controller"
	"github.com/hawks-atlanta/authentication-go/models"
	"gorm.io/gorm"
)

// GET /logs
func (r *Router) GetLogs(ctx *gin.Context) {

	logs, err := r.C.GetLogs()

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	ctx.JSON(http.StatusOK, logs)

}

// GET /logs/username
func (r *Router) GetLogByUser(ctx *gin.Context) {

	var filter controller.Filter[models.User]
	err := ctx.Bind(&filter)
	if err != nil {
		return
	}
	logs, err := r.C.GetLogsByUser(&filter)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // User was not found
			ctx.JSON(http.StatusNotFound, InternalServerError(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, logs)
}

// GET /logs/date
func (r *Router) GetLogByDate(ctx *gin.Context) {

	var filter controller.Filter[time.Time]
	err := ctx.Bind(&filter)
	if err != nil {
		return
	}

	logs, err := r.C.GetLogsByDate(&filter)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, InternalServerError(err))
	}

	ctx.JSON(http.StatusOK, logs)
}
