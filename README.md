# GoMetrics

GoMetrics is a lightweight server performance tool written in Go. it collects and analyzes key server metrics.

# Features

* Metrics collected:

  CPU usage

  Memory usage

  Disk usage

  Network I/O (Bytes Sent/Received)

  Disk I/O (Read/Write)

  Process-specific CPU and Memory usage

  System uptime

* Functionality

  Log metrics to `gometrics.go`

  Export metrics to `metrics.json`

  Write metrics to INFLUXDB for historical analysis

  Send slack alerts when threshold hits

# Installation

1. Clone the Repository:

   > git clone https://github.com/yourusername/GoMetrics.git

   > cd GoMetrics

2. Install Dependencies:

   > go mod tidy

3. Run the Project:

   > go run main.go

# Configuration

* InfluxDB:

  Set your InfluxDB token, organisation, and bucket in `utils/influxdb.go`

  Use enviroment variables for sensitive data:

  > export INFLUXDB_TOKEN="your-token"

* Slack Alerts:

  Update the Slack webhook URL in `utils/alert.go`

# Usage

Run the tool, and it will:

  * Print metrics to the console every 5 seconds.
  * Log metrics to `gometrics.log`.
  * Export metrics to `metrics.json`.
  * Write metrics to InfluxDB(if configured).
  * Send slack alerts if threshold are exceeded.

# Contributing

contribitions are welcome. Open an issue or submit a pull request.
