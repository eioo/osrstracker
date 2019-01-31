package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// PlayerStat on highscores
type PlayerStat struct {
	Level int64
	XP    int64
	Rank  int64
}

// PlayerStats map of all stats
type PlayerStats = map[string]PlayerStat

// GetHighscores fetch player highscores
func GetHighscores(playerName string) PlayerStats {
	stats := make(PlayerStats)
	skillNames := [24]string{
		"attack", "defence", "strength",
		"hitpoints", "ranged", "prayer",
		"magic", "cooking", "woodcutting",
		"fletching", "fishing", "firemaking",
		"crafting", "smithing", "mining",
		"herblore", "agility", "thieving",
		"slayer", "farming", "runecraft",
		"hunter", "construction",
	}
	resp, err := http.Get("https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws?player=" + playerName)

	if err != nil {
		// we are banned
	}

	scanner := bufio.NewScanner(resp.Body)
	i := 0

	for scanner.Scan() {
		str := strings.Trim(scanner.Text(), " \n")
		tokens := strings.Split(str, ",")

		if len(tokens) != 3 {
			continue
		}

		rank, _ := strconv.ParseInt(tokens[0], 10, 64)
		level, _ := strconv.ParseInt(tokens[1], 10, 64)
		xp, _ := strconv.ParseInt(tokens[2], 10, 64)

		skillName := skillNames[i]
		stat := PlayerStat{Rank: rank, Level: level, XP: xp}
		stats[skillName] = stat

		i++
	}

	fmt.Printf("%+v\n", stats)
	return stats
}

func main() {
	GetHighscores("aran penaali")
}
