package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type Data struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	for {
		resp, err := http.Get("http://localhost:8080/update")
		if err != nil {
			slog.Error("Error:", err)
			return
		}
		defer resp.Body.Close()

		var data Data
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			slog.Error("Error:", err)
			return
		}

		// log water, wind, and status
		slog.Debug(fmt.Sprintf("Water: %d meters, Wind: %d meters/sec, Status: %s\n", data.Water, data.Wind, data.Status))

		time.Sleep(15 * time.Second)
	}
}
