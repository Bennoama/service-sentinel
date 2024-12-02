package server

import (
	"log"
	"net/http"
	"service-sentinel/db"

	"github.com/gin-gonic/gin"
)

func GetWithGinContext (ctx *gin.Context) {
	monitors, err := db.GetAllMonitors()
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, monitors)
}

func Init() {
	router := gin.Default()
	router.GET("/monitors", GetWithGinContext)
	router.Run("localhost:8080")
}