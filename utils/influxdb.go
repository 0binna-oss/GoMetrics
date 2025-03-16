package utils

import (
	"log"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func WriteToInfluxDB(cpuUsage, memoryUsage, diskUsage float64, bytesSent, bytesRecv, uptime uint64) {
	token := os.Getenv("INFLUXDB_TOKEN")
	if token == "" {
		log.Fatal("InfluxDB token not found in enviroment variables")
	}

	client := influxdb2.NewClient("https://us-east-1-1.aws.cloud2.influxdata.com/orgs/c3082627a0aad62a", "ReNmHzMwvrI01q2VUEacmwCDCo1Wpy7Spo68eBnKkkLRfybIql8Q6riytUp8pFNNfuM1YIdX8ERXmYMAXihHcg==")

	writeAPI := client.WriteAPI("rich", "GoMetrics Bucket")
	p := influxdb2.NewPoint("server_metrics",
		map[string]string{},
		map[string]interface{}{
			"cpu":        cpuUsage,
			"memory":     memoryUsage,
			"disk":       diskUsage,
			"bytes_sent": bytesSent,
			"bytes_recv": bytesRecv,
			"uptime":     uptime,
		},
		time.Now())
	writeAPI.WritePoint(p)
	writeAPI.Flush()
	client.Close()
	log.Println("Metrics written to influxDB")
}
