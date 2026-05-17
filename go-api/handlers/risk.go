package handlers

import (
	"encoding/json"
	"net/http"
)

func RiskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := map[string]any{
		"ip":   "127.0.0.1",
		"risk": 82,
		"vpn":  true,
	}
	json.NewEncoder(w).Encode(data)
}
