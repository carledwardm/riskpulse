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
			{
				"ip":      "45.22.11.234",
				"country": "BR",
				"risk":    88,
				"vpn":     true,
				"status":  "flagged",
			},
			{
				"ip":      "103.44.52.12",
				"country": "JP",
				"risk":    12,
				"vpn":     false,
				"status":  "clean",
			},
			{
				"ip":      "185.122.3.44",
				"country": "RU",
				"risk":    75,
				"vpn":     true,
				"status":  "review",
			},
			{
				"ip":      "92.11.44.101",
				"country": "FR",
				"risk":    22,
				"vpn":     false,
				"status":  "clean",
			},
			{
				"ip":      "210.33.156.8",
				"country": "CN",
				"risk":    94,
				"vpn":     true,
				"status":  "flagged",
			},
			{
				"ip":      "88.201.3.15",
				"country": "IT",
				"risk":    41,
				"vpn":     false,
				"status":  "review",
			},
			{
				"ip":      "151.101.1.1",
				"country": "US",
				"risk":    5,
				"vpn":     false,
				"status":  "clean",
			},
			{
				"ip":      "77.242.10.33",
				"country": "NL",
				"risk":    61,
				"vpn":     true,
				"status":  "review",
			},
			{
				"ip":      "203.0.113.45",
				"country": "AU",
				"risk":    82,
				"vpn":     true,
				"status":  "flagged",
			},
		},
	}
	json.NewEncoder(w).Encode(data)
}
