package metrics

import (
	"github.com/shirou/gopsutil/v3/disk"
)

func GetDiskUsage() float64 {
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return -1
	}
	return diskUsage.UsedPercent
}
