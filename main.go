package main

import (
	"fmt"
	"gometrics/metrics"
	"gometrics/utils"
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

		// Display metrics
		fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
		fmt.Printf("Memory Usage: %.2f%%\n", memoryUsage)
		fmt.Printf("Disk Usage: %.2f%%\n", diskUsage)
		fmt.Printf("Network Usage: sent: %v bytes, Received: %v bytes\n", bytesSent, bytesRecv)
		fmt.Printf("Uptime: %v seconds\n", uptime)
		fmt.Println("-------------")

		// Log metrics
		utils.LogMetrics(cpuUsage, memoryUsage, diskUsage, bytesSent, bytesRecv, uptime)

		// Export metrics to JSON
		utils.ExportMetricsToJSON(cpuUsage, memoryUsage, diskUsage, bytesSent, bytesRecv, uptime)

		//alert sent if CPU usage is high
		if cpuUsage > 90 {
			utils.SendSlackAlert(fmt.Sprintf("High CPU usage detected: %.2f%%", cpuUsage))
		}

		time.Sleep(5 * time.Second)
	}
}
