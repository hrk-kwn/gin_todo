package app

import (
	"log"
	"github.com/gin-gonic/gin"
)

var (
	// https://github.com/gin-gonic/gin#quick-start
	// ここミドルウェア
	r = gin.Default()
)

// StartApp - Start our application
func StartApp() {
	mapUrls()
	log.Println("Start App...")

	r.Run(":8080")
}