package main

import (
	"github.com/adrien3d/adventofcode/utils"
	"strconv"
	"strings"
)

func solve(input string, v bool) (part1TotalScore, part2TotalScore int) {
	fishes := strings.Split(input, ",")
	var listP1 []int
	listP2 := make(map[int]int)

	for _, fish := range fishes {
		fishAge, _ := strconv.Atoi(fish)
		listP1 = append(listP1, fishAge)
		listP2[fishAge]++
	}

	for i := 0; i < 80; i++ {
		for j := range listP1 {
			if listP1[j] > 0 {
				listP1[j]--
			} else {
				listP1[j] = 6
				listP1 = append(listP1, 8)
			}
		}
	}
	for i := 0; i < 256; i++ {
		dying := listP2[0]
		for j := 0; j <= 8; j++ {
			listP2[j] = listP2[j+1]
		}
		listP2[6] += dying
		listP2[8] = dying
	}
	for i := 0; i <= 8; i++ {
		part2TotalScore += listP2[i]
	}

	return len(listP1), part2TotalScore
}

func main() {
	testInput := "3,4,3,1,2"
	realInput := "5,1,1,1,3,5,1,1,1,1,5,3,1,1,3,1,1,1,4,1,1,1,1,1,2,4,3,4,1,5,3,4,1,1,5,1,2,1,1,2,1,1,2,1,1,4,2,3,2,1,4,1,1,4,2,1,4,5,5,1,1,1,1,1,2,1,1,1,2,1,5,5,1,1,4,4,5,1,1,1,3,1,5,1,2,1,5,1,4,1,3,2,4,2,1,1,4,1,1,1,1,4,1,1,1,1,1,3,5,4,1,1,3,1,1,1,2,1,1,1,1,5,1,1,1,4,1,4,1,1,1,1,1,2,1,1,5,1,2,1,1,2,1,1,2,4,1,1,5,1,3,4,1,2,4,1,1,1,1,1,4,1,1,4,2,2,1,5,1,4,1,1,5,1,1,5,5,1,1,1,1,1,5,2,1,3,3,1,1,1,3,2,4,5,1,2,1,5,1,4,1,5,1,1,1,1,1,1,4,3,1,1,3,3,1,4,5,1,1,4,1,4,3,4,1,1,1,2,2,1,2,5,1,1,3,5,2,1,1,1,1,1,1,1,4,4,1,5,4,1,1,1,1,1,2,1,2,1,5,1,1,3,1,1,1,1,1,1,1,1,1,1,2,1,3,1,5,3,3,1,1,2,4,4,1,1,2,1,1,3,1,1,1,1,2,3,4,1,1,2"
	utils.Run(solve, testInput, true)
	utils.Run(solve, realInput, true)
}
