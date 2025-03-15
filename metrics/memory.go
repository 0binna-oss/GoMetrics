package metrics

import (
	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemoryUsage() float64 {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return -1
	}
	return memInfo.UsedPercent
}
