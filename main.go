package main

import (
	"log"
	"os"

	"github.com/caarlos0/env/v9"
	"github.com/gin-gonic/gin"
	"github.com/hawks-atlanta/authentication-go/controller"
	"github.com/hawks-atlanta/authentication-go/database"
	"github.com/hawks-atlanta/authentication-go/internal/config"
	"github.com/hawks-atlanta/authentication-go/router"
)

var environ config.Environment

func init() {
	gin.SetMode(gin.ReleaseMode)

	err := env.Parse(&environ)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := database.NewDB(environ.DatabaseEngine, environ.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	c, err := controller.New(controller.WithDB(db), controller.WithSecret(environ.JWTSecret))
	if err != nil {
		log.Fatal(err)
	}
	e := router.New(
		router.WithEngine(gin.Default()),
		router.WithController(c),
	)
	err = e.Run(os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
	}
}
