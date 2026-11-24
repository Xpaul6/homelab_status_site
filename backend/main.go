package main

import (
	"dash_backend/handlers"
	"dash_backend/services"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	VERSION = "2.1-k8s"
	PORT    = "8080"
)

var sysInfoMutex sync.Mutex

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	go services.StartInfoFetching(&sysInfoMutex)

	router := gin.Default()
	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "up"})
	})
	router.GET("/version", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, VERSION)
	})

	// Routes
	router.GET("/server-status", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Server is up!")
	})

	router.GET("/mcstatus/:address", handlers.MCStatusHandler())

	router.GET("/websitestatus", handlers.WebsiteStatusHandler())

	router.GET("/sysinfo", handlers.SysInfoHandler(&sysInfoMutex))

	router.Run(fmt.Sprintf(":%s", PORT))
}

