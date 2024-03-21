package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Data struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	StatusWater string `json:"statusWater"`
	StatusWind  string `json:"statusWind"`
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
		slog.Debug(fmt.Sprintf("Water: %d meters, Wind: %d meters/sec, Status Water: %s, Status Wind: %s\n", data.Water, data.Wind, data.StatusWater, data.StatusWind))

		createFile(data.Water, data.Wind)

		time.Sleep(15 * time.Second)
	}
}

func createFile(water, wind int) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.Create("file-output/status.json")
	if err != nil {
		log.Fatal(err)
	}

	jsonString := fmt.Sprintf(`
	{
		"status": {
			"water": %v,
			"wind": %v
		}
	}`, water, wind)

	_, err = f.Write([]byte(jsonString))

	if err != nil {
		log.Fatal("update json error: ", err.Error())
		return
	}
}
