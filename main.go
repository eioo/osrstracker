package main

import (
	"flag"
	"fmt"
)


func trackPlayer(playerName *string) {
	fmt.Println("Tracking OSRS player:", *playerName)
	stats := GetHighscores(*playerName)
	display(stats)
}

func main() {
	trackPtr := flag.String("track", "", "track player")
	flag.Parse()

	if *trackPtr != "" {
		trackPlayer(trackPtr)
	}
}
