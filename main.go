package main

import (
	"flag"
	"fmt"
)


func trackPlayer(playerName *string) {
	fmt.Println("Tracking OSRS player:", *playerName)
	stats := GetHighscores(*playerName)

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
