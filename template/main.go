package main

import (
	"github.com/adrien3d/adventofcode/2020/utils"
	"strings"
)

func solve(input string, v bool) (part1TotalScore, part2TotalScore int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		utils.Log(v, "info", line)
	}

	return part1TotalScore, part2TotalScore
}

func main() {
	testInput := ""
	realInput := ""
	utils.Run(solve, testInput, true)
	utils.Run(solve, realInput, true)
}
