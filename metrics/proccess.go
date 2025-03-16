package metrics

import (
	"github.com/shirou/gopsutil/v3/process"
)

// ProcessMetrics holds CPU and memory usage
type ProcessMetrics struct {
	CPUUsage    float64
	MemoryUsage float64
}

// GetProcessMetrics returns CPU and memory usage
func GetProcessMetrics(pid int32) (ProcessMetrics, error) {
	p, err := process.NewProcess(pid)
	if err != nil {
		return ProcessMetrics{}, err
	}
	cpuPercent, err := p.CPUPercent()
	if err != nil {
		return ProcessMetrics{}, err
	}
	memPercent, err := p.MemoryPercent()
	if err != nil {
		return ProcessMetrics{}, err
	}
	return ProcessMetrics{
		CPUUsage:    cpuPercent,
		MemoryUsage: float64(memPercent),
	}, nil
}
