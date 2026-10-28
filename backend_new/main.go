package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/sensors"
)

type MCInfo struct {
	Online  bool `json:"online"`
	Players struct {
		Online int `json:"online"`
		Max    int `json:"max"`
	} `json:"players"`
}

type SysInfo struct {
	CPU_temp         float64 `json:"cpu_temp"`
	CPU_load_percent float64 `json:"cpu_load_percent"`
	RAM_percent      float64 `json:"ram_percent"`
}

const (
	PORT           = "8080"
	CHECK_INTERVAL = 1500 * time.Millisecond
)

func main() {
	router := gin.Default()
	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "up"})
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
		body := MCInfo{}
		if err := json.Unmarshal(bodyBytes, &body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Incompatible data structure"})
			return
		}
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
		info, err := GetSysInfo()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get system metrics"})
			return
		}
		ctx.JSON(http.StatusOK, info)
	})

	router.Run(fmt.Sprintf(":%s", PORT))
}

func GetSysInfo() (SysInfo, error) {
	cpu_load_percent, err := cpu.Percent(CHECK_INTERVAL, false)
	if err != nil {
		return SysInfo{}, nil
	}

	tempSensors, err := sensors.SensorsTemperatures()
	var cpu_temp float64
	if err != nil {
		return SysInfo{}, nil
	}

	for _, v := range tempSensors {
		if v.SensorKey == "coretemp_package_id_0" || v.SensorKey == "PMU tcal" {
			cpu_temp = v.Temperature
			break
		}
	}

	vm, err := mem.VirtualMemory()
	if err != nil {
		return SysInfo{}, err
	}
	ram_percent := vm.UsedPercent
	res := SysInfo{
		CPU_temp:         cpu_temp,
		CPU_load_percent: cpu_load_percent[0],
		RAM_percent:      ram_percent,
	}
	return res, nil
}
