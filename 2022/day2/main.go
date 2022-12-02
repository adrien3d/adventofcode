package main

import (
	"github.com/adrien3d/adventofcode/utils"
	"strings"
)

func solve(input string, v bool) (part1TotalScore, part2TotalScore int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		roundScore, planScore, opponentScore := 0, 0, 0
		bets := strings.Split(line, " ")
		switch bets[0] {
		case "A":
			bets[0] = "rock"
			opponentScore += 1
		case "B":
			bets[0] = "paper"
			opponentScore += 2
		case "C":
			bets[0] = "scissors"
			opponentScore += 3
		}
		switch bets[1] {
		case "X": //loose
			bets[1] = "rock"
			roundScore += 1
			switch bets[0] {
			case "rock":
				planScore += 3
			case "paper":
				planScore += 1
			case "scissors":
				planScore += 2
			}
		case "Y": //draw
			bets[1] = "paper"
			roundScore += 2
			planScore += 3 + opponentScore
		case "Z": //win
			bets[1] = "scissors"
			roundScore += 3
			planScore += 6
			switch bets[0] {
			case "rock":
				planScore += 2
			case "paper":
				planScore += 3
			case "scissors":
				planScore += 1
			}
		}
		if bets[0] == bets[1] {
			roundScore += 3
		} else if ((bets[1] == "rock") && (bets[0] == "scissors")) || ((bets[1] == "paper") && (bets[0] == "rock")) || ((bets[1] == "scissors") && (bets[0] == "paper")) {
			roundScore += 6
		}

		utils.Log(v, "info", bets, roundScore, planScore)
		part1TotalScore += roundScore
		part2TotalScore += planScore
	}

	return part1TotalScore, part2TotalScore
}

func main() {
	testInput := "A Y\nB X\nC Z"
	realInput := "A Z\nC X\nA Z\nA Z\nC Y\nC Y\nA Z\nA Y\nC Y\nA Y\nA Z\nA Z\nA Z\nA Y\nA Z\nA Y\nC Y\nC X\nA Y\nC Y\nC Y\nC X\nA Z\nC Y\nC X\nA X\nA Y\nA Z\nA Y\nA Y\nC X\nC X\nB Y\nC X\nC X\nA Y\nA Z\nA Z\nA X\nA Z\nA Z\nC Y\nA Z\nA Z\nA Y\nC X\nC Y\nC X\nB X\nC Z\nA Y\nA Z\nA Z\nA Z\nA Y\nA Y\nC X\nA Y\nA Z\nC Y\nA Y\nA Y\nA Z\nA Z\nC Y\nA Z\nC Y\nA Y\nA Z\nA Z\nC Y\nB Z\nA Z\nA Z\nA Z\nA Z\nC X\nC X\nA Y\nA X\nA Y\nA Z\nA Z\nC X\nA Z\nC X\nA Z\nC Y\nC X\nA Z\nA Z\nA Z\nA Z\nC X\nC Y\nA Z\nA Z\nC Y\nA Z\nB Z\nC X\nA Z\nA X\nC X\nA Z\nC Z\nA Z\nC X\nA Z\nA Y\nC X\nC Z\nA Z\nC X\nC Y\nA Z\nB Z\nB Y\nA X\nA X\nA X\nA Z\nA Z\nA X\nA Z\nA X\nA Z\nA Z\nC X\nC X\nB Z\nA Z\nA Y\nA Z\nA Z\nA Z\nA Z\nC X\nC X\nC X\nA Z\nA Z\nA Z\nA Z\nA X\nA Y\nA Y\nC X\nC X\nB Z\nC X\nA X\nA Z\nA Z\nC X\nC Z\nA Z\nA Z\nC Y\nA X\nA Z\nC Y\nA X\nA Y\nA Y\nA Y\nC Y\nA Z\nA Z\nC X\nC X\nC X\nC X\nB Z\nC Y\nC X\nC X\nA Y\nA Y\nA Z\nA Z\nA X\nC X\nA Z\nB Z\nA Z\nC X\nA Y\nA Z\nA Z\nA Y\nA Y\nA Z\nC X\nA Y\nC X\nC X\nA Z\nA Y\nC X\nA Z\nA Z\nA Y\nA Z\nA Z\nC Y\nC Z\nA Y\nA Z\nC X\nC X\nA Z\nA Z\nC Y\nA Y\nA Z\nA Y\nA Z\nC X\nA Z\nA Z\nC X\nA Y\nA Y\nC Y\nA Z\nC Y\nA Z\nA Z\nC Y\nA Y\nB Z\nC Y\nC X\nC X\nA Z\nC Y\nA X\nC Y\nA Z\nA Y\nA Z\nC X\nC X\nC Y\nC Y\nA Y\nA Z\nC X\nC X\nA Y\nA X\nA Y\nB Z\nA Y\nC Y\nC X\nC X\nA X\nC X\nB Z\nA X\nC Y\nC X\nC X\nA X\nA Z\nB Z\nA Z\nA Z\nA Y\nA X\nA Z\nC X\nA X\nC Y\nA Z\nA Z\nA X\nA Z\nA Z\nC X\nC X\nA Z\nA Z\nA Z\nA X\nA Z\nA X\nB Y\nA Z\nA Y\nC Y\nA Z\nC X\nA Z\nC X\nA Y\nA Z\nC X\nC Y\nA Y\nC Y\nA Z\nC X\nA Z\nA Z\nC X\nA Y\nA Z\nA X\nA Z\nA Z\nA Z\nB Y\nC X\nA X\nA Z\nA Z\nC Y\nC Y\nC X\nC X\nC X\nA X\nA X\nA Z\nA X\nA Y\nA Z\nA Y\nA Z\nC Y\nC Z\nA Y\nA Z\nA Z\nA X\nA Z\nA Z\nC X\nC Y\nA Y\nC X\nC X\nA Z\nC X\nC X\nC X\nA X\nA Z\nA Z\nA Z\nA Z\nB Y\nA Y\nA Y\nA Z\nC X\nA Y\nA Z\nC X\nA Z\nC Z\nA Y\nC X\nA Z\nB Z\nA Y\nA X\nA Z\nC X\nA Z\nA Z\nA Z\nA Z\nA Z\nB Z\nA Z\nC X\nA Y\nC X\nA Z\nA Z\nA Y\nA Z\nB Z\nC X\nA Y\nC Y\nA Z\nA Z\nC X\nC X\nA Y\nC X\nC Y\nB Z\nA Y\nC X\nA Y\nC X\nA X\nA Y\nA Z\nA Z\nA Y\nC X\nA X\nC X\nB Y\nA Z\nA Y\nB Y\nA Y\nC X\nA Z\nA Z\nC X\nC Y\nA Y\nC X\nC Y\nA Y\nA Z\nA X\nB Z\nC X\nA Z\nA Y\nA Z\nB Z\nA Z\nA X\nC Y\nA X\nA Z\nA Y\nC Y\nA Z\nC Y\nA Z\nC X\nC X\nA Y\nC X\nC X\nA Y\nA Z\nA Z\nA Y\nA X\nC Y\nA Z\nA Z\nC X\nA X\nA Z\nC Z\nA Z\nC Y\nA Z\nC X\nA Z\nA Z\nA X\nC X\nC X\nC Y\nB Z\nB Y\nC Y\nA Y\nA X\nA Z\nC X\nA Y\nA Y\nA Z\nA Z\nC Z\nC X\nC X\nC X\nA Z\nC X\nA Z\nA Z\nA Y\nC Y\nC X\nC X\nC X\nA Y\nC X\nB Z\nC X\nA Z\nC Y\nA Y\nC Y\nA Z\nA Z\nC X\nA X\nA Z\nA Y\nA Z\nB X\nC Z\nA Z\nC Z\nA X\nC X\nC X\nC X\nC X\nA Z\nA Y\nA Z\nA Z\nA Z\nA Z\nA Z\nC Z\nC Y\nC Y\nC X\nC Y\nA Z\nC X\nC X\nA Z\nA Z\nC X\nC Y\nC Y\nA Z\nA X\nC Y\nC Y\nC X\nA Y\nC X\nA Z\nA Z\nA Y\nC Y\nA Y\nC X\nC X\nA Z\nA Z\nC Y\nA Z\nC X\nA Y\nA Y\nA Z\nC Y\nA Z\nC X\nA Y\nA Z\nA Z\nA Z\nA Z\nC Y\nA Y\nC Y\nA Z\nA Z\nA X\nA Y\nA Y\nA Z\nC Y\nA Z\nA X\nA Z\nB Z\nC X\nC X\nC Y\nA Z\nA Z\nC X\nC Z\nA Z\nC X\nC Y\nA Z\nB Z\nA Z\nB Z\nA X\nA Y\nA Z\nA Z\nA Z\nA Z\nC Y\nA Z\nA Z\nA Z\nB Z\nA Z\nC X\nC X\nA Z\nC X\nA X\nA Z\nA Y\nA Y\nA X\nA Z\nA Z\nA Z\nA Z\nB Z\nA Z\nC Y\nC Y\nC X\nC X\nB Z\nC Y\nA Z\nC Y\nA Z\nA Y\nC Y\nA X\nA Y\nC X\nA Y\nC X\nA Z\nA Z\nB Z\nA Y\nC Y\nC X\nA Z\nA Z\nC X\nA Y\nA Z\nA Z\nA X\nA Y\nA Y\nA Y\nA Z\nA Y\nA Z\nC X\nA X\nA Z\nA Z\nC Y\nA Z\nC X\nA Z\nC Y\nA Y\nA Z\nA Z\nA Z\nA Z\nC Y\nA X\nA Z\nA Z\nA Y\nA Z\nB Z\nA X\nA Y\nC X\nC X\nA Y\nA Z\nC Y\nA Z\nA Y\nA Y\nA Z\nA Y\nA Y\nA Z\nA Z\nA Z\nC X\nA Z\nA Y\nA Z\nA Z\nC Y\nA Z\nC X\nA Y\nC X\nC Y\nA Z\nC Y\nA Z\nA Y\nC Y\nC Y\nA Z\nC X\nC X\nC X\nA Z\nA Z\nA Y\nC X\nA Z\nA X\nA Z\nA Z\nC X\nA Z\nC Y\nA Y\nA Z\nA Y\nA Z\nA Z\nC Y\nA Z\nC X\nA Z\nA X\nA Z\nA Z\nC X\nA Z\nA Z\nC Y\nC X\nA Z\nC Y\nC X\nC X\nA Z\nA Z\nA Z\nA X\nC X\nA Z\nA Z\nC Y\nA Y\nC X\nA Y\nC X\nC Y\nA Y\nA Z\nC X\nA Z\nB Y\nA X\nB Z\nA Y\nA X\nC Y\nA Z\nA X\nA Z\nA Z\nC Y\nA Z\nB Z\nC X\nC X\nB Y\nA Z\nA Z\nA Z\nC Y\nC X\nA Z\nA Y\nA Y\nA Y\nC Y\nC X\nA Z\nA X\nA X\nA X\nA Z\nA Z\nA Z\nA Z\nC X\nC X\nA Z\nC X\nA Y\nC X\nB Z\nA Z\nA Z\nC X\nA Z\nC X\nC X\nC X\nA Y\nC Y\nC X\nA Y\nC X\nC Y\nB Z\nA Z\nC X\nA Z\nA X\nB Z\nA Y\nB Y\nA Z\nA Z\nA X\nA Z\nA X\nA Z\nA Z\nC X\nC Y\nA X\nC Y\nC X\nA Z\nA Z\nA Z\nC Y\nA X\nA Y\nA Z\nA Y\nC X\nB Z\nA Z\nA Y\nC Z\nC X\nA Z\nA Z\nA Z\nB Z\nA X\nC X\nA Z\nA Z\nB Z\nA Z\nA Z\nB Z\nA Z\nC X\nA Z\nC X\nA Z\nC X\nA Z\nC Y\nA Z\nA X\nA Y\nA Y\nC X\nA Y\nC X\nB Z\nA Z\nA Z\nA Z\nA Z\nC X\nA Z\nC X\nA Z\nA Y\nA Z\nB Z\nA Y\nC Z\nA Y\nC X\nA Z\nA Z\nA Y\nB Z\nA X\nC Y\nA Z\nA Y\nA Z\nA Y\nA Y\nA Z\nA Z\nA Z\nA X\nA Z\nC Y\nA Y\nA X\nA Y\nC X\nA Y\nC X\nA X\nC X\nA Z\nC Y\nA Y\nA Z\nA Z\nA X\nA Y\nC X\nC X\nA Y\nA X\nA Z\nB Z\nA Y\nA Z\nA X\nA Z\nB Y\nA Y\nA Y\nA Z\nC X\nA Z\nA Z\nA Y\nA Z\nC Z\nA Z\nA Z\nA Z\nA Z\nC Y\nC Y\nA Y\nB Z\nC Y\nA Y\nC Y\nA Z\nA Z\nA Z\nA Z\nC Y\nA Z\nA Z\nC Y\nC X\nA Y\nA Y\nA Z\nC X\nC Z\nC X\nC X\nA Z\nA Z\nA Y\nA Z\nA X\nC Y\nA Z\nA Z\nC Y\nC X\nA X\nA Z\nA Z\nA X\nC X\nC X\nC X\nA Z\nA X\nC X\nC Y\nA X\nA Z\nC X\nA Z\nA Z\nC X\nA Y\nA Z\nA Z\nA Y\nA Z\nC X\nA Z\nA Z\nC X\nA Y\nA Z\nA Z\nC X\nA Z\nA Y\nA Z\nC Y\nA X\nA Z\nA Z\nC X\nA Z\nA Y\nC Y\nB Z\nA Z\nA Y\nC X\nA Z\nB Z\nA Z\nC Z\nA X\nA Z\nA Z\nC Y\nA Z\nA Y\nC X\nC Y\nA Z\nA Z\nA Y\nA X\nC Y\nA Y\nC X\nC Y\nA Z\nC Y\nA Z\nC X\nA Z\nA Z\nA X\nA Z\nA Z\nB X\nA X\nA Z\nC Y\nA Z\nA X\nC X\nA Z\nA Z\nC X\nC Y\nC Y\nA X\nA Y\nC Y\nA Y\nA Z\nA Z\nA Z\nA Z\nA Y\nC X\nC X\nC Y\nC X\nA Z\nA X\nB Z\nB Y\nC X\nC Y\nA Y\nA Z\nA Y\nC X\nC Z\nA Z\nA Y\nC Y\nC X\nA Z\nA Z\nA Y\nC X\nC Z\nC Y\nA Z\nC X\nC Y\nA X\nA X\nA Y\nA Z\nB Z\nA X\nA Y\nA Y\nC X\nC Y\nA Z\nA X\nA Z\nA X\nA Y\nA Z\nA Z\nA Z\nC X\nA Z\nA Z\nA Z\nC Z\nC Y\nC Y\nA Z\nC Y\nC Y\nC Y\nC X\nA Z\nC X\nC X\nA Z\nA Y\nA Z\nA Z\nA X\nA Y\nA Y\nC X\nC X\nA Z\nA Z\nA Z\nA Z\nA Y\nA Z\nA Z\nA Z\nA Z\nA X\nA X\nA Y\nA X\nC Y\nA Y\nA Z\nC X\nA Y\nA Y\nA Z\nA Z\nA Z\nC X\nA Z\nC X\nC X\nC Y\nA Y\nA Z\nA Y\nA Z\nC X\nC X\nA Z\nA Z\nA Z\nC X\nA X\nA Z\nA Z\nC Y\nC Y\nA Y\nA Y\nA Z\nA Z\nC Y\nC X\nC Y\nA X\nC Y\nC X\nC Y\nA Z\nA Z\nA X\nC Y\nC Y\nA Z\nA Y\nC X\nA X\nB Z\nA Z\nC X\nA Y\nA X\nA Z\nA Z\nA Z\nC Y\nC X\nA Z\nC X\nA Z\nA Z\nA Y\nA Z\nA Y\nA X\nA Z\nC Z\nA Z\nA Z\nA Z\nA X\nA Z\nA X\nC X\nA Z\nA Z\nC Y\nA X\nA Z\nC X\nC X\nA Y\nA Z\nA Z\nC X\nB X\nA Z\nC Y\nA Z\nC X\nA Y\nA Z\nC X\nA X\nA Z\nA X\nA Z\nC X\nA Z\nA X\nA Z\nC X\nC Y\nA Y\nA X\nA X\nA Z\nC Y\nC Y\nA Y\nA Y\nA Y\nA X\nA Z\nA Z\nA Z\nC Y\nC Y\nA Y\nB X\nB X\nA Z\nC X\nC X\nA Z\nC Y\nC X\nA Z\nA Z\nA Z\nA Z\nC X\nA Z\nA Z\nC Y\nA Y\nC Y\nA Z\nC Y\nC Y\nA Z\nC X\nA Z\nA X\nA Z\nC X\nC X\nA Y\nB Z\nA Y\nC Y\nA Z\nC Y\nA X\nC X\nA Y\nC X\nA Z\nC Z\nC Y\nA Z\nC X\nC Y\nA Z\nA X\nA Z\nA Z\nA Z\nC X\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nA Y\nA Z\nA X\nA Y\nC Y\nB Y\nC X\nB Z\nA Z\nA Z\nA Y\nB Z\nA Z\nA Z\nC X\nC Y\nC X\nA Y\nA X\nC X\nC Y\nA Y\nC X\nC Y\nA Z\nA Z\nA Z\nA Y\nC X\nA Y\nC X\nB Z\nA Z\nA Y\nA Z\nA Z\nA Y\nC X\nA Z\nC X\nC Y\nA Y\nA Z\nB Z\nC X\nA Y\nC X\nA Y\nA Z\nC X\nA X\nC X\nB Y\nC X\nA Z\nA Y\nA Z\nA Y\nA X\nC X\nC X\nA Y\nC X\nA Y\nA Y\nA X\nB Y\nA Y\nC X\nC X\nA Y\nB Z\nB X\nB Z\nA Y\nA Z\nC Y\nA Y\nB Y\nA Z\nC X\nA Z\nA Z\nA Z\nA Z\nB Z\nC X\nC Y\nA Z\nC Y\nC Y\nA X\nC X\nA Z\nA Z\nA Z\nC X\nC X\nC X\nA Y\nC Z\nC Z\nA Z\nC X\nA Y\nA Z\nA Z\nA Z\nC X\nA Z\nA X\nA Z\nA Z\nA Z\nA Z\nA Y\nC Y\nC X\nA X\nA Y\nC X\nA X\nA Z\nA Z\nC X\nA Z\nA X\nA Z\nA Z\nA X\nA Z\nA Z\nA X\nA Z\nA X\nB Y\nA Y\nA Y\nC Y\nA Z\nA Y\nC X\nA Z\nA Y\nA X\nC Y\nB X\nC Y\nA Z\nC X\nA Y\nA Z\nA Y\nA X\nC Y\nA Z\nA Z\nC Y\nC X\nA Z\nC X\nA Y\nC X\nA Z\nA Y\nA Z\nA Z\nA Z\nA Y\nA Z\nC X\nC X\nA X\nC X\nC X\nA Z\nC X\nA Z\nC Y\nC X\nA Z\nA Z\nA Z\nC X\nA X\nC Y\nA Z\nC Y\nA X\nA Z\nC X\nA Z\nA Z\nA X\nA Z\nC X\nB X\nA Z\nA Z\nA Z\nC X\nA Y\nA Y\nA X\nC Y\nC Y\nA Z\nA Y\nA Z\nA Z\nC X\nA X\nA Y\nA Z\nA Z\nA Z\nA Z\nB Z\nC X\nC X\nC X\nA Z\nC Z\nA X\nC X\nA Z\nC Y\nA Z\nA Z\nA Y\nA Y\nC X\nA Z\nA X\nA Z\nA Z\nA Z\nA Z\nC Y\nA Z\nA Y\nA X\nA X\nA Z\nC X\nA X\nA X\nA Z\nA Y\nC X\nA Z\nA Z\nA Y\nA Z\nB Z\nC X\nC X\nC Z\nC Y\nC X\nA Z\nC Y\nA Z\nC Z\nA Z\nA Y\nA Y\nA X\nA X\nA Z\nA Y\nA Y\nA Y\nA Y\nA Z\nC Y\nA Z\nA Z\nC X\nA Z\nA Z\nC Y\nA Y\nC X\nA Y\nC X\nA Z\nB Z\nA X\nB X\nA Y\nA X\nA Y\nB Z\nA Y\nA Z\nC Y\nC Y\nA Z\nA X\nA Z\nA Z\nC Z\nA Z\nA Y\nC X\nA Y\nC X\nA X\nA Y\nC Y\nA Y\nA Z\nA Z\nC X\nC X\nB Z\nA Z\nA Z\nA X\nC X\nC Y\nA Z\nA Z\nA X\nC X\nC Z\nA Z\nC Y\nA Y\nB Z\nC Y\nA Z\nC X\nA X\nA Z\nA Z\nA Z\nA Y\nC Y\nA Z\nC Y\nA Z\nA Z\nA Z\nA Z\nA Y\nC X\nA Y\nC Y\nB Z\nA Z\nC X\nC Y\nA Z\nC X\nA Z\nC X\nC Z\nA Z\nC X\nC X\nA Z\nA Z\nA Y\nA Y\nA Y\nC X\nA Y\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nA Y\nA Z\nA X\nC Y\nA Z\nA Z\nA Z\nA Y\nA Z\nA Z\nA Z\nA Z\nC X\nB Z\nA Z\nA Y\nA Y\nA X\nA Z\nA Z\nC Z\nA Z\nC X\nA Y\nA X\nB Z\nA Z\nA Z\nA Z\nC Y\nC Z\nC X\nA Z\nC Y\nC Y\nC Y\nC X\nB Z\nA Z\nA Z\nC Z\nA X\nA Z\nA Z\nA Z\nC X\nA Z\nA Z\nC Y\nC Y\nA Z\nA Z\nC X\nA Y\nC Y\nC Y\nA Z\nA Z\nA X\nA Z\nA Z\nA X\nA X\nC X\nA Z\nA X\nC X\nC X\nA Z\nA X\nA Z\nC X\nC X\nC X\nC Y\nA Z\nA Z\nA X\nA Z\nA Y\nA Z\nA Z\nC Y\nA Z\nA Z\nA Z\nA X\nA Z\nC X\nA X\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nA Y\nA X\nA Y\nA X\nC Y\nA Z\nA Z\nC X\nA Z\nA Z\nA Z\nA X\nA Z\nA Z\nA Y\nC X\nA Y\nA Z\nC X\nA X\nA Y\nA Z\nA Z\nA X\nA Z\nA Y\nC X\nA Y\nA Z\nA Z\nA Z\nA Z\nC X\nC Y\nA Z\nB Z\nC X\nA Z\nA Z\nC Y\nC Y\nC X\nA X\nC Y\nB Y\nA Z\nA Z\nA Z\nC Y\nA Z\nA X\nA Y\nA Z\nA Z\nA Z\nC Y\nA Z\nC Y\nC X\nA Z\nA Z\nA Z\nA Z\nA Z\nA X\nA Z\nA Z\nC X\nA Z\nA Y\nC Z\nA Z\nA Z\nA Z\nB X\nC X\nA Z\nA Z\nA Z\nA Z\nC X\nA Z\nA Z\nA Y\nA X\nC X\nC Y\nA X\nA Y\nC X\nA Z\nA Z\nC X\nC X\nA Z\nC X\nA Z\nA Y\nC X\nA Z\nB X\nB Y\nA X\nC Y\nA X\nA Y\nC Y\nA Z\nA Z\nA Z\nC X\nA Z\nA X\nA Z\nC Y\nA Z\nA Z\nC X\nA Z\nA Z\nA Y\nA Z\nA X\nA Y\nA Z\nC X\nC Y\nB X\nC Y\nA Y\nA Z\nC X\nA Z\nC X\nC Y\nA X\nA X\nA Z\nA Z\nC Y\nA Y\nA Y\nA Z\nA Z\nC X\nA X\nC X\nA X\nA Y\nC X\nA Z\nA Z\nB Y\nA Z\nA Z\nA Z\nA Z\nA Z\nA X\nA Y\nA X\nA Z\nA Z\nC X\nA Y\nA Z\nC Y\nC X\nC X\nC X\nC X\nC Y\nA X\nA X\nC Y\nA X\nA Y\nA Y\nB Z\nA Z\nB Y\nC Y\nA X\nA Y\nA Z\nA Z\nA Z\nA X\nC X\nA Z\nA Z\nA Z\nA Z\nA Z\nA Y\nA Z\nB Z\nA Z\nA Z\nA Y\nC Y\nC Y\nC X\nA Z\nA X\nC X\nA Y\nB Z\nC X\nA Z\nC X\nC Y\nC Y\nA X\nA X\nC X\nA Z\nA Z\nA Y\nA X\nA X\nA Z\nC Y\nB X\nA Z\nA Z\nA Z\nA Z\nA Y\nA Y\nA Z\nC Y\nC Y\nA X\nA Z\nA Z\nC Y\nA Y\nA Y\nA X\nA Y\nA X\nA X\nA X\nA Z\nA Z\nA Y\nA Z\nC X\nA Z\nA Y\nA Z\nA Z\nA Y\nA Z\nA Z\nC X\nA Z\nA Y\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nA Y\nC X\nB Z\nC Z\nA Z\nA Y\nA X\nA X\nC Y\nC Y\nA X\nC X\nB Z\nA Z\nC X\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nA Y\nA Z\nA Z\nA Y\nA Z\nA Y\nA Y\nA Z\nA Z\nA Y\nA Z\nC Y\nA Z\nA Z\nA Z\nC X\nC X\nA Z\nC Y\nA Z\nA Z\nC X\nA Z\nC X\nA Y\nA Z\nC Y\nA Z\nA Z\nC X\nC Y\nA Z\nC X\nC X\nA X\nA Y\nC Y\nA Y\nA Z\nC X\nA Z\nA Z\nA X\nA Z\nC Y\nA Z\nC X\nA Z\nC X\nA Z\nC X\nA Y\nC Y\nA X\nA Z\nA Z\nC Y\nA X\nA Z\nA Y\nB Z\nA Z\nA X\nA Z\nA Z\nA X\nC X\nA Z\nA Z\nC Y\nA Z\nA Y\nA Z\nC Y\nA Z\nA Y\nA Z\nC Z\nA Y\nA Z\nA Y\nC Y\nA Z\nC X\nA X\nB Z\nC X\nC X\nA Z\nA Z\nA Y\nA X\nA Z\nA Z\nB Z\nA Z\nC Z\nA X\nA Z\nA Z\nA Z\nA Z\nA X\nA X\nA Z\nA X\nA Z\nA Z\nA Z\nA X\nC X\nC Y\nA X\nC X\nC X\nA Z\nA X\nC Y\nB Z\nA Z\nA Z\nC X\nB X\nA Z\nA Z\nC X\nC X\nB Z\nC Y\nA Z\nA Z\nC X\nA X\nA Y\nC Y\nC Y\nB Z\nA Y\nC X\nA Z\nA Z\nA Y\nC X\nA Z\nA Z\nC X\nC Y\nA Z\nC X\nC Y\nC Y\nC X\nC Y\nA Z\nC Y\nA Z\nC X\nA Z\nC Y\nC Z\nA Z\nB Z\nA Z\nC X\nC X\nB Y\nB Z\nC Y\nC X\nC X\nA Y\nC X\nC X\nA Z\nA Y\nA Z\nC X\nA X\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nC X\nC X\nA Z\nC Y\nC X\nA Z\nB Z\nA Z\nA Z\nA X\nC X\nA Y\nA Z\nA Z\nA Y\nA X\nC Y\nB X\nA Z\nA X\nC Y\nC X\nC X\nC Y\nC Y\nA Z\nA Z\nC Y\nA Z\nA Y\nC Y\nA Y\nC X\nA Z\nC X\nC Y\nC Y\nA Z\nA Z\nA X\nA Y\nA Z\nA X\nA Z\nA Z\nA Z\nC Y\nA Z\nA Z\nC X\nC Y\nA Z\nA Z\nA Z\nC Y\nC X\nA Z\nC X\nA Z\nA Z\nA Z\nA X\nA Z\nA Z\nA Y\nB Z\nA Z\nA Z\nA Z\nC Y\nA Z\nB Z\nA Z\nC Z\nA Z\nA Y\nC X\nC Y\nC X\nC X\nA Z\nA Z\nA Y\nA Y\nA Z\nA Y\nB Z\nC Y\nA Y\nA Z\nC X\nA Z\nA Z\nC Y\nA Y\nA Z\nA Y\nC Y\nA Z\nA Z\nA X\nB Z\nA Z\nA X\nC X\nA Z\nC Y\nC Y\nA Z\nB Y\nA Y\nA Z\nA Z\nA Z\nA Z\nC X\nC X\nA Z\nC X\nA Y\nA Z\nA Z\nC Y\nA Z\nA Z\nC Y\nC X\nB Z\nA Y\nA Y\nC X\nC X\nA Z\nA X\nB Z\nA Z\nC Y\nA Z\nA Y\nA Z\nA Z\nA Y\nC Y\nC X\nA Z\nC X\nA Z\nC Y\nA Z\nC Y\nA Y\nA Z\nA Y\nA Z\nC X\nA Z\nA Z\nC X\nA Z\nB Y\nA Z\nA Z\nC Y\nC X\nC X\nA Z\nA Z\nC X\nB Z\nA Y\nA Z\nA Y\nA Z\nA Z\nA Y\nA X\nC X\nC X\nA X\nA Z\nA Y\nA Y\nA Z\nA Y\nA Z\nA Z\nC Y\nA X\nA Z\nA Z\nC X\nA Z\nA X\nB X\nC X\nA Z\nA Y\nB Z\nC X\nC Y\nA Z\nB Z\nC Y\nA Z\nA Z\nA X\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z\nA Z"
	utils.Run(solve, testInput, true)
	utils.Run(solve, realInput, true) //12830 too high, 10000 too low
}
