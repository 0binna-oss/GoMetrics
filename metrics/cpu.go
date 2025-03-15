package metrics

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCPUUsage() float64 {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return -1
	}
	return cpuPercent[0]
}
