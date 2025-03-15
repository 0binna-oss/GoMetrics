package utils

import (
	"encoding/json"
	"os"
	"time"
)

type Metrics struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage   float64 `json:"disk_usage"`
	BytesSent   uint64  `json:"bytes_sent"`
	BytesRecv   uint64  `json:"bytes_recv"`
	Uptime      uint64  `json:"uptime"`
	Timestamp   string  `json:"timestamp"`
}

func ExportMetricsToJSON(cpuUsage, memoryUsage, diskUsage float64, bytesSent, bytesRecv, uptime uint64) {
	metrics := Metrics{
		CPUUsage:    cpuUsage,
		MemoryUsage: memoryUsage,
		DiskUsage:   diskUsage,
		BytesSent:   bytesSent,
		BytesRecv:   bytesRecv,
		Uptime:      uptime,
		Timestamp:   time.Now().Format(time.RFC3339),
	}
	jsonData, err := json.MarshalIndent(metrics, "", " ")
	if err != nil {
		return
	}
	os.WriteFile("metrics.json", jsonData, 0644)
}
