package main

import (
	"net/http"
	"riskpulse/go-api/handlers"
)

func main() {
	http.HandleFunc("/risk", handlers.RiskHandler)
	http.ListenAndServe(":8080", nil)
}
