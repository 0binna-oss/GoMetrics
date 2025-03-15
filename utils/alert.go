package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendSlackAlert(message string) {
	webhookURL := "https://hooks.slack.com/services/your/webhook/url"
	payload := map[string]string{"text": message}
	jsonData, _ := json.Marshal(payload)

	_, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Failed to send Slack alert:", err)
	}
}
