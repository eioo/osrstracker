package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

// GetHighscores fetch player highscores
func GetHighscores(playerName string) {
	resp, err := http.Get("https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws?player=" + playerName)

	if err != nil {
		// we are banned
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		str := strings.Trim(scanner.Text(), " \n")
		tokens := strings.Split(str, ",")

		if len(tokens) != 3 {
			continue
		}

		fmt.Printf(scanner.Text() + "\n")
	}
}

func main() {
	GetHighscores("aran penaali")
}
