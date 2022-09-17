package server

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

var router *gin.Engine

func Run() {
	router = gin.Default()

	router.LoadHTMLGlob("")

	router.Run()
}
