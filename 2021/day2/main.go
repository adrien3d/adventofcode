package main

import (
	"fmt"
	"strconv"
	"strings"
)

type instruct struct {
	Move string
	Step int
}

func solvePart1(input string) (horPos, depth int) {
	lines := strings.Split(input, "\n")
	list := make([]instruct, len(lines))
	for i, line := range lines {
		lin := strings.Split(line, " ")
		st, _ := strconv.Atoi(lin[1])
		list[i] = instruct{
			Move: lin[0],
			Step: st,
		}
		if lin[0] == "forward" {
			horPos += st
		} else if lin[0] == "up" {
			depth -= st
		} else if lin[0] == "down" {
			depth += st
		}
	}
	fmt.Println(horPos * depth)
	return horPos, depth
}

func solvePart2(input string) (horPos, depth int) {
	lines := strings.Split(input, "\n")
	list := make([]instruct, len(lines))
	var aim int
	for i, line := range lines {
		lin := strings.Split(line, " ")
		st, _ := strconv.Atoi(lin[1])
		list[i] = instruct{
			Move: lin[0],
			Step: st,
		}
		if lin[0] == "forward" {
			depth += aim * st
			horPos += st
		} else if lin[0] == "up" {
			aim -= st
		} else if lin[0] == "down" {
			aim += st
		}
	}
	fmt.Println(horPos * depth)
	return horPos, depth
}

func main() {
	testInput := "forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2"
	part1Input := "forward 8\ndown 6\ndown 8\nforward 7\ndown 5\nup 2\ndown 3\ndown 7\ndown 8\ndown 8\ndown 8\ndown 2\nup 1\ndown 3\nup 2\ndown 4\ndown 2\nforward 6\nforward 4\ndown 3\ndown 2\nforward 2\nforward 1\nforward 4\nforward 5\nforward 8\ndown 1\ndown 4\nup 5\nup 2\nforward 3\ndown 9\nforward 7\nforward 9\nforward 9\nforward 8\ndown 1\ndown 2\nforward 7\ndown 3\nforward 6\ndown 4\nforward 7\ndown 1\nup 8\nforward 3\ndown 1\nforward 7\nup 1\nforward 8\nup 6\nup 2\ndown 6\nforward 1\nup 6\nforward 5\ndown 9\nup 5\nforward 7\nforward 9\ndown 9\ndown 3\nforward 7\nforward 8\nforward 3\nforward 9\nforward 7\ndown 3\ndown 7\ndown 4\nforward 2\ndown 7\ndown 3\ndown 5\nup 1\ndown 9\nup 4\nforward 1\nup 9\ndown 2\nforward 8\ndown 8\ndown 6\nforward 7\ndown 9\ndown 3\nforward 8\nforward 3\ndown 6\ndown 7\ndown 4\nforward 3\ndown 3\ndown 9\nforward 8\nforward 9\nup 5\nforward 1\ndown 3\ndown 3\ndown 3\ndown 9\ndown 2\ndown 9\nforward 5\nup 3\nup 5\nup 7\ndown 2\ndown 7\ndown 9\ndown 5\ndown 4\ndown 8\nforward 1\nup 8\nup 3\nforward 1\nforward 5\nforward 3\nup 7\ndown 9\ndown 9\nforward 7\ndown 1\nforward 1\nforward 8\nforward 6\ndown 1\ndown 7\nforward 9\nup 4\nforward 8\nup 6\nforward 3\ndown 3\ndown 9\nforward 5\nup 3\ndown 7\nforward 9\nforward 2\nup 1\nforward 7\nup 8\nforward 7\nforward 1\nup 3\nup 7\ndown 1\nforward 5\nup 8\ndown 2\nup 2\nup 3\ndown 5\nforward 6\nup 8\ndown 7\nup 8\nup 4\ndown 8\nforward 9\ndown 8\ndown 2\nup 7\ndown 5\nforward 1\nup 1\ndown 1\nforward 1\nforward 1\nforward 3\nforward 8\ndown 4\ndown 5\nforward 9\nup 6\nup 7\ndown 8\nforward 8\ndown 2\nforward 6\ndown 3\nforward 9\nforward 5\nup 7\ndown 2\nup 6\nup 6\ndown 9\nforward 3\nup 1\nup 2\nforward 9\ndown 1\nup 3\nforward 4\nforward 9\ndown 3\ndown 4\nforward 4\nup 6\nup 5\nforward 2\ndown 5\ndown 1\nforward 9\ndown 7\nup 6\nup 5\nforward 4\nforward 9\ndown 6\nforward 1\nup 6\ndown 1\nforward 4\nup 9\ndown 6\nforward 5\ndown 2\nforward 8\nforward 9\ndown 7\ndown 4\ndown 1\nforward 1\ndown 4\ndown 6\nforward 5\nforward 2\nforward 8\nforward 5\ndown 6\nup 9\nforward 2\ndown 1\nforward 6\nforward 6\ndown 5\nforward 5\ndown 8\nforward 3\ndown 5\nup 1\nforward 4\ndown 5\ndown 4\nforward 4\ndown 3\ndown 5\ndown 7\nforward 5\nforward 2\nup 2\nup 4\nforward 7\ndown 3\ndown 1\ndown 7\nup 8\nforward 6\nforward 3\nforward 7\nforward 5\nup 5\ndown 3\ndown 6\nforward 7\nup 9\nup 5\nforward 2\ndown 9\nforward 8\nforward 6\nforward 5\nup 5\ndown 9\ndown 8\nup 2\nup 4\nforward 5\nforward 2\nup 4\nforward 3\ndown 7\nforward 8\nforward 1\nforward 9\nforward 6\nup 7\nup 2\nforward 1\ndown 5\nforward 9\ndown 8\ndown 4\ndown 7\nup 2\ndown 5\nforward 7\nup 3\nforward 6\ndown 2\nforward 8\nforward 8\nup 3\nforward 6\nforward 9\nforward 8\nforward 3\nup 9\nforward 9\ndown 6\nforward 5\nforward 8\nup 1\nforward 2\nforward 6\nforward 8\nup 6\ndown 3\ndown 9\ndown 6\nup 7\nforward 6\nforward 1\nforward 1\nforward 7\ndown 5\ndown 9\ndown 3\nup 3\nforward 3\nforward 2\ndown 5\nup 4\nforward 1\ndown 9\nforward 9\nforward 1\nforward 1\ndown 9\ndown 2\nforward 4\nforward 9\ndown 5\nup 5\ndown 6\nforward 8\ndown 4\ndown 1\nup 5\nup 3\ndown 2\ndown 3\nforward 8\nforward 5\nforward 9\ndown 4\nup 9\ndown 1\nforward 2\ndown 8\nup 2\ndown 8\nup 6\nforward 7\ndown 1\nup 7\ndown 9\nforward 9\ndown 9\nforward 7\nforward 4\ndown 5\nup 3\ndown 3\nforward 8\ndown 3\ndown 4\ndown 9\nforward 4\nup 4\nforward 6\ndown 1\nforward 5\ndown 2\nforward 6\ndown 4\ndown 1\nforward 3\nup 3\nup 3\nforward 8\nforward 6\nforward 6\ndown 9\nforward 5\ndown 9\nforward 6\nforward 3\nup 4\nforward 6\ndown 8\nup 3\ndown 9\ndown 3\nforward 6\ndown 4\ndown 8\ndown 6\ndown 5\nforward 1\ndown 3\nforward 9\ndown 9\ndown 3\nforward 9\ndown 2\nforward 3\nup 6\nforward 2\nforward 1\nforward 8\ndown 2\ndown 2\ndown 7\nup 7\nforward 3\nup 2\nup 6\nup 6\ndown 2\nforward 2\nforward 2\ndown 6\ndown 2\nup 6\nforward 4\ndown 9\nup 3\ndown 4\nforward 7\nup 6\nforward 3\nforward 1\ndown 1\ndown 8\ndown 8\ndown 1\nforward 2\ndown 6\ndown 6\nforward 2\nup 6\ndown 2\nup 4\ndown 1\nup 8\nup 5\ndown 4\nforward 2\nforward 2\ndown 2\nforward 9\ndown 5\ndown 9\nforward 6\ndown 9\ndown 5\ndown 7\ndown 3\nup 9\ndown 6\nup 6\nup 8\nforward 8\nforward 8\ndown 3\nup 9\nforward 9\nforward 8\nforward 6\ndown 4\ndown 6\nup 9\ndown 9\ndown 5\nup 2\nup 2\nforward 2\nforward 1\ndown 5\ndown 8\nup 3\nforward 2\ndown 1\ndown 9\nforward 7\nforward 5\nup 3\nup 6\ndown 5\nup 1\ndown 2\nup 7\nforward 1\ndown 6\nup 6\nup 1\nup 2\nforward 2\ndown 4\nup 1\nup 3\nup 9\nup 7\nforward 4\ndown 5\ndown 9\ndown 8\nforward 1\ndown 4\nforward 4\nforward 8\nup 4\ndown 8\ndown 1\ndown 9\ndown 5\nforward 3\nforward 8\nup 2\ndown 6\nup 6\nforward 5\ndown 6\ndown 8\nforward 6\ndown 6\nup 5\ndown 2\nup 5\ndown 7\ndown 9\nforward 3\ndown 8\nforward 1\nforward 5\nforward 2\ndown 4\nforward 2\nforward 7\nup 7\nup 3\ndown 2\nforward 7\nup 6\nforward 6\nforward 1\ndown 4\ndown 2\ndown 6\ndown 1\nforward 1\nforward 8\ndown 1\nup 2\ndown 2\ndown 1\ndown 6\nforward 7\nforward 6\nforward 5\ndown 1\ndown 8\ndown 1\nup 5\nforward 6\nforward 5\nup 5\nforward 5\nup 8\ndown 3\nforward 1\nforward 6\nup 8\nup 9\ndown 7\ndown 1\nforward 2\nforward 1\nforward 9\nforward 3\nforward 7\nforward 8\ndown 6\nup 5\ndown 1\nforward 1\nforward 8\ndown 6\nforward 7\nforward 8\ndown 7\ndown 5\ndown 7\nup 7\ndown 5\nforward 5\ndown 4\ndown 7\nforward 6\nforward 5\nforward 6\nforward 7\nup 9\ndown 2\ndown 2\ndown 4\ndown 8\nup 3\ndown 7\ndown 5\nforward 6\ndown 9\ndown 5\ndown 9\ndown 1\nforward 6\nup 7\ndown 2\ndown 2\nforward 8\nforward 1\ndown 3\ndown 4\nforward 3\nforward 4\ndown 1\nforward 9\nup 7\nforward 8\ndown 9\nforward 7\nforward 6\nforward 2\ndown 8\nup 9\ndown 2\nforward 8\nup 7\ndown 5\ndown 9\ndown 3\ndown 6\ndown 4\nup 2\ndown 3\ndown 1\nup 1\nup 6\nforward 4\ndown 1\nforward 1\nup 4\nforward 4\nforward 3\nforward 8\nforward 9\nforward 9\ndown 2\ndown 5\nup 8\nup 1\ndown 9\nforward 5\ndown 1\nup 5\ndown 4\nup 3\nforward 9\nup 7\nforward 9\nup 1\nforward 4\nforward 8\nup 6\ndown 6\ndown 8\ndown 8\ndown 9\ndown 2\nup 7\nforward 9\nup 8\ndown 9\nup 6\nforward 4\nup 7\ndown 6\nup 7\ndown 4\nforward 2\nforward 9\ndown 6\ndown 8\nforward 6\nforward 3\ndown 3\nforward 3\nforward 7\nup 2\ndown 8\nforward 7\ndown 5\ndown 1\ndown 6\ndown 5\ndown 2\nup 6\nforward 7\nforward 6\ndown 1\ndown 5\nforward 7\nforward 3\ndown 9\ndown 8\nforward 5\nup 7\nforward 1\nup 5\ndown 7\nforward 8\nforward 6\nforward 2\ndown 1\ndown 9\nup 1\ndown 2\ndown 2\ndown 7\ndown 4\nforward 1\ndown 3\ndown 5\nup 8\nforward 7\nup 5\ndown 8\ndown 6\ndown 3\ndown 3\ndown 9\ndown 7\nforward 4\nup 5\nforward 3\nforward 7\ndown 3\nup 6\nforward 4\nforward 4\ndown 4\ndown 2\nup 1\nforward 8\nforward 3\nup 1\nforward 1\ndown 9\ndown 6\nup 1\ndown 4\ndown 8\nup 9\nforward 2\ndown 3\nforward 8\ndown 6\ndown 5\ndown 4\nup 5\ndown 9\nup 3\nforward 4\ndown 9\ndown 7\nforward 6\nforward 6\nforward 8\nforward 6\ndown 9\ndown 1\nforward 3\nforward 9\nforward 4\nup 8\nup 5\nup 2\ndown 9\nforward 9\nforward 3\nforward 5\nup 8\ndown 2\ndown 1\nforward 9\nforward 7\ndown 7\nforward 1\ndown 5\ndown 8\ndown 4\ndown 7\ndown 1\ndown 4\ndown 7\nforward 2\ndown 5\nforward 1\ndown 4\ndown 5\ndown 2\nup 5\nforward 9\ndown 5\nforward 1\ndown 7\ndown 4\ndown 7\ndown 6\nforward 5\ndown 3\ndown 1\nup 2\nforward 2\nforward 2\nforward 1\ndown 1\nforward 3\nforward 5\nforward 4\ndown 7\nforward 7\ndown 1\nforward 7\nforward 5\ndown 8\nforward 6\nforward 6\nforward 6\nforward 7\nup 9\ndown 4\ndown 1\ndown 8\nforward 7\nup 4\nforward 4\ndown 6\nup 1\nforward 5\nforward 2\ndown 1\nforward 7\nforward 6\nforward 5\nforward 2\ndown 5\ndown 6\ndown 9\nup 4\nforward 6\nforward 2\ndown 5\ndown 3\nup 4\ndown 6\nup 8\nforward 8\nup 9\nforward 6\nforward 6\nup 5\ndown 7\nforward 9\nforward 6\ndown 9\ndown 9\nup 1\nforward 7\ndown 6\nup 4\ndown 8\ndown 3\nforward 9\nforward 5\nforward 9\ndown 2\nforward 3\ndown 1\nforward 9\nup 4\nup 8\nforward 6\ndown 1\nforward 9\nforward 4\ndown 5\nforward 2\nup 3\nforward 5\nup 8\nup 7\ndown 8\nforward 4\ndown 6\nforward 7\nup 2\ndown 2\nforward 4\ndown 9\ndown 8\nforward 2\nforward 2\ndown 2\ndown 3\nforward 3\ndown 1\nforward 8\ndown 7\nup 9\ndown 4\ndown 2\ndown 5\nup 7\ndown 8\ndown 2\ndown 4\ndown 4\ndown 8\nforward 7\nforward 7\ndown 8\nup 2\nup 3\nforward 8\nup 1\ndown 7\nforward 7\ndown 6\ndown 8\nup 6\nforward 5\nforward 3\ndown 6\nforward 9\nup 4\nup 7\nforward 4\ndown 1\ndown 8\ndown 1\nforward 9\ndown 3\nforward 8\nforward 6\nforward 4\ndown 9\nforward 3\nup 5\nup 8\ndown 9\ndown 5\ndown 1\nup 8\nforward 8\nup 6\nforward 2\ndown 8\nup 4\nup 7\nforward 7\nforward 5\nforward 9\nforward 2\nup 4\ndown 9\nforward 7\ndown 6\ndown 6\nforward 7\ndown 5\nup 6\ndown 9\nforward 3"

	fmt.Println(solvePart1(testInput))
	fmt.Println(solvePart1(part1Input))

	fmt.Println(solvePart2(testInput))
	fmt.Println(solvePart2(part1Input))
}
