package utils

import "fmt"

// CheckThresholds check if metrics exceed custom thresholds and sends alerts
func CheckThresholds(cpuUsage, memoryUsage, diskUsage float64) {
	if cpuUsage > 80 {
		SendSlackAlert(fmt.Sprintf("High CPU usage detected: %.2f%%", cpuUsage))
	}
	if memoryUsage > 90 {
		SendSlackAlert(fmt.Sprintf("High memory usage detected: %.2f%%", memoryUsage))
	}
	if diskUsage > 85 {
		SendSlackAlert(fmt.Sprintf("High disk usage detected: %.2f%%", diskUsage))
	}
}
