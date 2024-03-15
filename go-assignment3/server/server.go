package main

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"net/http"
)

type Status string

const (
	Aman   Status = "aman"
	Siaga  Status = "siaga"
	Bahaya Status = "bahaya"
)

type Data struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status Status `json:"status"`
}

func main() {
	http.HandleFunc("/update", updateHandler)
	slog.Info("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		water int = rand.Intn(100)
		wind  int = rand.Intn(100)
	)

	data := Data{
		Water:  water,
		Wind:   wind,
		Status: getStatus(water, wind),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

	slog.Info("func updateHandler called")
}

func getStatus(water, wind int) Status {
	if water < 5 || wind < 6 {
		return Aman
	} else if (water >= 6 && water <= 8) || (wind >= 7 && wind <= 15) {
		return Siaga
	} else {
		return Bahaya
	}
}
