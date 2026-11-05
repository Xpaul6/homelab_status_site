package main

import (
	"dash_backend/models"
	"dash_backend/services"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	VERSION        = "2.1-k8s"
	PORT           = "8080"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
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

	router.GET("/mcstatus/:address", func(ctx *gin.Context) {
		address := ctx.Param("address")

		resp, err := http.Get(fmt.Sprintf("https://api.mcstatus.io/v2/status/java/%s", address))
		if err != nil {
			log.Print(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch server status"})
			return
		}
		if resp.StatusCode == http.StatusServiceUnavailable {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Data API unavalible"})
			return
		}

		bodyBytes, _ := io.ReadAll(resp.Body)
		body := models.MCInfo{}
		if err := json.Unmarshal(bodyBytes, &body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Incompatible data structure"})
			return
		}
		resp.Body.Close()
		ctx.JSON(http.StatusOK, body)
	})

	router.GET("/websitestatus", func(ctx *gin.Context) {
		address := ctx.Query("address")
		_, err := http.Get(address)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": fmt.Sprintf("Cannot reach %s", address)})
			return
		}
		ctx.Status(http.StatusOK)
	})

	router.GET("/sysinfo", func(ctx *gin.Context) {
		info, err := services.GetSysInfo()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get system metrics"})
			return
		}
		ctx.JSON(http.StatusOK, info)
	})

	router.Run(fmt.Sprintf(":%s", PORT))
}

