package metrics

import (
	"github.com/shirou/gopsutil/v3/net"
)

func GetNetworkUsage() (uint64, uint64) {
	stats, err := net.IOCounters(false)
	if err != nil {
		return 0, 0
	}
	return stats[0].BytesSent, stats[0].BytesRecv
}
