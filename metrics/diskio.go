package metrics

import (
	"github.com/shirou/gopsutil/v3/disk"
)

// returns total read and write bytes across all disk
func GetDiskIO() (uint64, uint64, error) {
	stats, err := disk.IOCounters()
	if err != nil {
		return 0, 0, err
	}
	var readBytes, writeBytes uint64
	for _, stat := range stats {
		readBytes += stat.ReadBytes
		writeBytes += stat.WriteBytes
	}
	return readBytes, writeBytes, nil
}
