package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// PlayerStat on high scores
type PlayerStat struct {
	Name  string `json:"name"`
	Level int64  `json:"lvl"`
	XP    int64  `json:"xp"`
	Rank  int64  `json:"rank"`
}

type PlayerStats = []PlayerStat

func (stat *PlayerStat) display()  {
	fmt.Printf("%-14vRank: %-12v XP: %-12v Level: %v\n",
				stat.Name, stat.Rank, stat.XP, stat.Level)
}

// GetHighscores fetch player highscores
func GetHighscores(playerName string) PlayerStats {
	var stats PlayerStats
	resp, err := http.Get("https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws?player=" + playerName)

	if resp == nil ||  err != nil {
		log.Printf("Failed to fetch highscore for \"%s\"", playerName)
		os.Exit(0)
	}

	scanner := bufio.NewScanner(resp.Body)
	i := 0

	for scanner.Scan() {
		str := strings.Trim(scanner.Text(), " \n")
		tokens := strings.Split(str, ",")
		tokenCount := len(tokens)

		if tokenCount != 3 {
			continue
		}

		rank, _ := strconv.ParseInt(tokens[0], 10, 64)
		level, _ := strconv.ParseInt(tokens[1], 10, 64)
		xp, _ := strconv.ParseInt(tokens[2], 10, 64)

		skillName := HiscoreResult[i]
		stat := PlayerStat{Name: skillName, Rank: rank, Level: level, XP: xp}
		stats = append(stats, stat)

		i++
	}

	return stats
}
