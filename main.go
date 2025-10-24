package main

import (
 "encoding/json"
 "log"
 "net/http"
 "sync"
 "time"
)

type Stats struct{ TotalTransfers int `json:"totalTransfers"` }
type Donation struct{ TxnID, Date string `json:"txnId","date"` }

var (
 mu     sync.Mutex
 stats  = Stats{TotalTransfers: 0}
 last   = Donation{}
)

func main(){
 http.HandleFunc("/stats", handleStats)
 http.HandleFunc("/last-donation", handleLast)
 http.HandleFunc("/simulate", handleSim)
 log.Println("Amina Water backend running on :8080")
 log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleStats(w http.ResponseWriter, _ *http.Request){
 mu.Lock(); defer mu.Unlock()
 json.NewEncoder(w).Encode(stats)
}
func handleLast(w http.ResponseWriter, _ *http.Request){
 mu.Lock(); defer mu.Unlock()
 json.NewEncoder(w).Encode(last)
}
func handleSim(w http.ResponseWriter, _ *http.Request){
 mu.Lock()
 stats.TotalTransfers++
 if stats.TotalTransfers>=1000 {
  last = Donation{TxnID:"SIMULATED_"+time.Now().Format("150405"),Date:time.Now().Format(time.RFC822)}
  stats.TotalTransfers=0
 }
 mu.Unlock()
 w.WriteHeader(204)
}
