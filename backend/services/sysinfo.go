package services

import (
	"time"

	"dash_backend/models"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/sensors"
)

const CHECK_INTERVAL = 1500 * time.Millisecond
 
func GetSysInfo() (models.SysInfo, error) {
	cpu_load_percent, err := cpu.Percent(CHECK_INTERVAL, false)
	if err != nil {
		return models.SysInfo{}, nil
	}

	tempSensors, err := sensors.SensorsTemperatures()
	var cpu_temp float64
	if err != nil {
		return models.SysInfo{}, nil
	}

	for _, v := range tempSensors {
		if v.SensorKey == "coretemp_package_id_0" || v.SensorKey == "PMU tcal" {
			cpu_temp = v.Temperature
			break
		}
	}

	vm, err := mem.VirtualMemory()
	if err != nil {
		return models.SysInfo{}, err
	}
	ram_percent := vm.UsedPercent
	res := models.SysInfo{
		CPU_temp:         cpu_temp,
		CPU_load_percent: cpu_load_percent[0],
		RAM_percent:      ram_percent,
	}
	return res, nil
}
