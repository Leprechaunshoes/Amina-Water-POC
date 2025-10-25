package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// --- Data Structures ---
type Stats struct {
	TotalTransfers int    `json:"totalTransfers"`
	LastDonation   string `json:"lastDonation"`
}

type Donation struct {
	TxnID string `json:"txnID"`
	Date  string `json:"date"`
}

// --- Global Variables ---
var (
	mu    sync.Mutex
	stats = Stats{TotalTransfers: 0, LastDonation: ""}
)

// --- Handlers ---
func handleStats(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(stats)
}

func handleSim(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	stats.TotalTransfers++
	if stats.TotalTransfers >= 1000 {
		stats.LastDonation = time.Now().Format("2006-01-02 15:04:05")
		log.Println("💧 Amina Humanity Fund: $100 donation triggered!")
		stats.TotalTransfers = 0
	}
	json.NewEncoder(w).Encode(stats)
}

// --- Main Server ---
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Amina Water backend live ✓ — use /stats or /simulate"))
	})
	http.HandleFunc("/stats", handleStats)
	http.HandleFunc("/simulate", handleSim)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🌍 Amina Water backend running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}