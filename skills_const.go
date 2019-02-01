package main

var HiscoreResult = map[int]string{
	0: "overall",
	1: "attack",
	2: "defence",
	3: "strength",
	4: "hitpoints",
	5: "ranged",
	6: "prayer",
	7: "magic",
	8: "cooking",
	9: "woodcutting",
	10: "fletching",
	11: "fishing",
	12: "firemaking",
	13: "crafting",
	14: "smithing",
	15: "mining",
	16: "herblore",
	17: "agility",
	18: "thieving",
	19: "slayer",
	20: "farming",
	21: "runecraft",
	22: "hunter",
	23: "construction",
	/*
	24: "bountyHunterHunter",
	25: "bountyHunterRogue",
	26: "clueScrollAll",
	27: "clueScrollEasy",
	28: "clueScrollMedium",
	29: "clueScrollHard",
	30: "clueScrollElite",
	31: "clueScrollMaster",
	32: "lastManStanding",
	*/
}

/*
func skillToId(skillName string) int {
	for i, val := range HiscoreResult {
		if val == skillName {
			return i
		}
	}

	return -1
}
*/