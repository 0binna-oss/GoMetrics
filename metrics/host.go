package metrics

import (
	"github.com/shirou/gopsutil/v3/host"
)

func GetUptime() uint64 {
	uptime, err := host.Uptime()
	if err != nil {
		return 0
	}
	return uptime
}
