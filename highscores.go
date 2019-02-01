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

// SkillStat on high scores
type SkillStat struct {
	Level int64
	XP    int64
	Rank  int64
}

// PlayerStats map of all stats
type PlayerStats = map[string]SkillStat

// GetHighscores fetch player highscores
func GetHighscores(playerName string) PlayerStats {
	stats := make(PlayerStats)
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
		stat := SkillStat{Rank: rank, Level: level, XP: xp}
		stats[skillName] = stat

		i++
	}

	return stats
}

func display(stats PlayerStats)  {
	for key := range stats {
		stat := stats[key]
		fmt.Printf("%-14vRank: %-12v XP: %-12v Level: %v\n", key, stat.Rank, stat.XP, stat.Level)
	}
}
