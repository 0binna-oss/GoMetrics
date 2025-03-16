package main

import (
	"fmt"
	"gometrics/metrics"
	"gometrics/utils"
	"log"
	"time"
)

func main() {
	fmt.Println("GoMetrics - Server Perfomance Analyzer")
	fmt.Println("============")

	for {
		// Collect metrics
		cpuUsage := metrics.GetCPUUsage()
		memoryUsage := metrics.GetMemoryUsage()
		diskUsage := metrics.GetDiskUsage()
		bytesSent, bytesRecv := metrics.GetNetworkUsage()
		uptime := metrics.GetUptime()
		readBytes, writeBytes, _ := metrics.GetDiskIO()
		processMetrics, err := metrics.GetProcessMetrics(1)

		if err != nil {
			log.Printf("Error getting process metrics: %v", err)
		}

		// Display metrics
		fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
		fmt.Printf("Memory Usage: %.2f%%\n", memoryUsage)
		fmt.Printf("Disk Usage: %.2f%%\n", diskUsage)
		fmt.Printf("Network Usage: sent: %v bytes, Received: %v bytes\n", bytesSent, bytesRecv)
		fmt.Printf("Disk I/O: Read: %v bytes, Write: %v bytes\n", readBytes, writeBytes)
		fmt.Printf("Process (PID 1) CPU: %.2f%%, Memory: %.2f%%\n", processMetrics.CPUUsage, processMetrics.MemoryUsage)
		fmt.Printf("Uptime: %v seconds\n", uptime)
		fmt.Println("-------------")

		// Log metrics
		utils.LogMetrics(cpuUsage, memoryUsage, diskUsage, bytesSent, bytesRecv, uptime, readBytes, writeBytes, processMetrics.CPUUsage, processMetrics.MemoryUsage)

		// Export metrics to JSON
		utils.ExportMetricsToJSON(cpuUsage, memoryUsage, diskUsage, bytesSent, bytesRecv, uptime)

		// Check thresholds and send alert
		utils.CheckThresholds(cpuUsage, memoryUsage, diskUsage)

		// Write metrics to influxDB
		utils.WriteToInfluxDB(cpuUsage, memoryUsage, diskUsage, bytesSent, bytesRecv, uptime)

		//alert sent if CPU usage is high
		if cpuUsage > 90 {
			utils.SendSlackAlert(fmt.Sprintf("High CPU usage detected: %.2f%%", cpuUsage))
		}

		time.Sleep(5 * time.Second)
	}
}
