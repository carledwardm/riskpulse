package handlers

import (
	"encoding/json"
	"net/http"
)

func RiskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data := map[string]any{
		"stats": map[string]any{
			"threat_requests": 128,
			"vpn_detections":  42,
			"fraud_attempts":  19,
			"risk_alerts":     7,
		},

		"risks": []map[string]any{
			{
				"ip":      "192.168.1.0",
				"country": "US",
				"risk":    91,
				"vpn":     true,
				"status":  "flagged",
			},
			{
				"ip":      "10.0.0.5",
				"country": "UK",
				"risk":    35,
				"vpn":     false,
				"status":  "clean",
			},
			{
				"ip":      "172.16.0.9",
				"country": "DE",
				"risk":    67,
				"vpn":     true,
				"status":  "review",
			},
		},
	}
	json.NewEncoder(w).Encode(data)
}
