package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func saveStats(playerName *string, stats *PlayerStats) {
	playerDir := "stats/" + *playerName
	createDirectory(playerDir)
	epoch := strconv.FormatInt(epochTime(), 10)
	logFile := fmt.Sprintf("%s/%s.log", playerDir, epoch)

	marshalled, err := json.Marshal(*stats)

	if err != nil {
		log.Printf("Failed to marshal stats")
		os.Exit(0)
	}

	writeToFile(logFile, &marshalled)
}

func trackPlayer(playerName *string) {
	fmt.Println("Tracking OSRS player:", *playerName)
	stats := GetHighscores(*playerName)

	if len(stats) == 0 {
		return
	}

	saveStats(playerName, &stats)

	// Debug
	for _, stat := range stats {
		stat.display()
	}
}

func main() {
	trackPtr := flag.String("track", "", "track player")
	flag.Parse()

	if *trackPtr != "" {
		trackPlayer(trackPtr)
	}
}
