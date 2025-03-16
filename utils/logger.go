package utils

import (
	"log"
	"os"
)

func LogMetrics(cpuUsage, memoryUsage, diskUsage float64, bytesSent, bytesRecv, uptime uint64, readBytes, writeBytes uint64, processCPU, processMem float64) {
	// Create or open the log file
	file, err := os.OpenFile("gometrics.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	logger.Printf("cpu Usage: %.2f%%, Memory Usage: %.2f%%, Disk Usage: %.2f%%\n, Network Sent: %v bytes, Network Received: %v bytes, Disk I/O Read: %v bytes, Disk I/O Write: %v bytes, Process CPU: %.2f%%, Process Memory: %.2f%%, uptime: %v seconds\n",
		cpuUsage, memoryUsage, diskUsage, bytesSent, bytesRecv, readBytes, writeBytes, processCPU, processMem, uptime)
}
