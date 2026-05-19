package main

import (
	"net/http"
	"os"
	"riskpulse/go-api/handlers"
)

func main() {
	http.HandleFunc("/risk", handlers.RiskHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
