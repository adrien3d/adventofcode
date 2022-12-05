package main

import (
	"container/list"
	"fmt"
	"github.com/adrien3d/adventofcode/utils"
	"strings"
)

type Action struct {
	Amount, From, To int
}

func printList(list *list.List) {
	v := list.Front()
	for i := 0; i < list.Len(); i++ {
		fmt.Print(string(v.Value.(int32)))
		v = v.Next()
	}
	fmt.Println()
}

func parse(input string) (stacks map[int]*list.List, actions []Action) {
	lines := strings.Split(input, "\n")
	stacks = make(map[int]*list.List)
	for _, line := range lines {
		if len(line) == 0 || line[1] == '1' {
			continue
		}
		if line[0] == ' ' || line[0] == '[' {
			for i, char := range line {
				if char == ' ' || char == '[' || char == ']' {
					continue
				}
				stackI := i/4 + 1
				stack, exist := stacks[stackI]
				if !exist {
					stack = list.New()
					stacks[stackI] = stack
				}
				stack.PushBack(char)
			}
		} else if line[0] == 'm' {
			action := Action{}
			fmt.Sscanf(line, "move %d from %d to %d", &action.Amount, &action.From, &action.To)
			actions = append(actions, action)
		}
	}

	return stacks, actions
}

func solve(input string, v bool) (part1TotalScore, part2TotalScore int) {
	stacks, actions := parse(input)
	stacksP2, _ := parse(input)
	/*for _, stack := range stacks {
		printList(stack)
	}
	fmt.Println(actions)*/

	for _, action := range actions {
		stackP2 := list.New()
		for i := 0; i < action.Amount; i++ {
			stacks[action.To].PushFront(stacks[action.From].Remove(stacks[action.From].Front()))
			stackP2.PushBack(stacksP2[action.From].Remove(stacksP2[action.From].Front()))
		}
		stacksP2[action.To].PushFrontList(stackP2)
	}

	word, wordP2 := []int32(nil), []int32(nil)
	for i := 1; i <= len(stacks); i++ {
		word = append(word, stacks[i].Front().Value.(int32))
	}
	for i := 1; i <= len(stacksP2); i++ {
		wordP2 = append(wordP2, stacksP2[i].Front().Value.(int32))
	}
	fmt.Println(string(word), string(wordP2))

	return part1TotalScore, part2TotalScore
}

func main() {
	testInput := "    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2"
	realInput := "        [F] [Q]         [Q]        \n[B]     [Q] [V] [D]     [S]        \n[S] [P] [T] [R] [M]     [D]        \n[J] [V] [W] [M] [F]     [J]     [J]\n[Z] [G] [S] [W] [N] [D] [R]     [T]\n[V] [M] [B] [G] [S] [C] [T] [V] [S]\n[D] [S] [L] [J] [L] [G] [G] [F] [R]\n[G] [Z] [C] [H] [C] [R] [H] [P] [D]\n 1   2   3   4   5   6   7   8   9 \n\nmove 3 from 5 to 2\nmove 3 from 8 to 4\nmove 7 from 7 to 3\nmove 14 from 3 to 9\nmove 8 from 4 to 1\nmove 1 from 7 to 5\nmove 2 from 6 to 4\nmove 4 from 5 to 7\nmove 1 from 3 to 6\nmove 3 from 4 to 3\nmove 1 from 4 to 1\nmove 5 from 1 to 9\nmove 1 from 4 to 6\nmove 4 from 7 to 4\nmove 15 from 9 to 2\nmove 7 from 1 to 6\nmove 3 from 3 to 5\nmove 1 from 4 to 9\nmove 2 from 5 to 3\nmove 2 from 4 to 9\nmove 4 from 1 to 6\nmove 1 from 3 to 1\nmove 1 from 3 to 2\nmove 4 from 6 to 3\nmove 24 from 2 to 8\nmove 4 from 9 to 8\nmove 1 from 1 to 3\nmove 2 from 5 to 4\nmove 1 from 2 to 4\nmove 19 from 8 to 1\nmove 5 from 3 to 9\nmove 8 from 1 to 3\nmove 3 from 4 to 1\nmove 6 from 9 to 5\nmove 2 from 3 to 4\nmove 1 from 8 to 5\nmove 2 from 4 to 6\nmove 11 from 6 to 1\nmove 8 from 8 to 7\nmove 1 from 6 to 5\nmove 13 from 1 to 3\nmove 1 from 1 to 7\nmove 2 from 7 to 8\nmove 5 from 7 to 1\nmove 2 from 8 to 4\nmove 3 from 5 to 3\nmove 11 from 3 to 1\nmove 2 from 5 to 3\nmove 2 from 5 to 3\nmove 2 from 7 to 1\nmove 7 from 3 to 1\nmove 1 from 4 to 5\nmove 1 from 6 to 4\nmove 3 from 4 to 7\nmove 3 from 7 to 1\nmove 6 from 3 to 5\nmove 1 from 5 to 9\nmove 4 from 5 to 4\nmove 2 from 3 to 4\nmove 8 from 9 to 2\nmove 5 from 4 to 6\nmove 1 from 6 to 5\nmove 1 from 4 to 9\nmove 39 from 1 to 7\nmove 7 from 2 to 6\nmove 1 from 9 to 3\nmove 1 from 2 to 7\nmove 1 from 3 to 1\nmove 5 from 7 to 3\nmove 4 from 5 to 1\nmove 19 from 7 to 9\nmove 1 from 9 to 8\nmove 1 from 9 to 7\nmove 5 from 9 to 3\nmove 6 from 6 to 7\nmove 1 from 8 to 3\nmove 4 from 1 to 4\nmove 23 from 7 to 6\nmove 1 from 1 to 6\nmove 21 from 6 to 2\nmove 3 from 4 to 8\nmove 7 from 6 to 1\nmove 1 from 4 to 9\nmove 1 from 6 to 7\nmove 6 from 1 to 2\nmove 1 from 7 to 4\nmove 15 from 2 to 8\nmove 5 from 3 to 8\nmove 22 from 8 to 7\nmove 1 from 8 to 1\nmove 5 from 3 to 4\nmove 1 from 3 to 2\nmove 1 from 1 to 2\nmove 3 from 4 to 8\nmove 3 from 8 to 9\nmove 11 from 2 to 1\nmove 2 from 1 to 4\nmove 15 from 9 to 5\nmove 22 from 7 to 3\nmove 2 from 4 to 9\nmove 3 from 4 to 2\nmove 8 from 1 to 8\nmove 6 from 8 to 6\nmove 1 from 6 to 2\nmove 3 from 6 to 9\nmove 3 from 2 to 7\nmove 4 from 2 to 9\nmove 2 from 7 to 5\nmove 1 from 1 to 7\nmove 2 from 8 to 2\nmove 2 from 7 to 5\nmove 9 from 5 to 3\nmove 8 from 5 to 2\nmove 1 from 6 to 4\nmove 1 from 6 to 9\nmove 1 from 2 to 9\nmove 2 from 5 to 1\nmove 7 from 2 to 3\nmove 1 from 4 to 3\nmove 1 from 2 to 4\nmove 5 from 3 to 4\nmove 6 from 9 to 3\nmove 1 from 2 to 6\nmove 6 from 9 to 6\nmove 2 from 1 to 8\nmove 3 from 6 to 3\nmove 2 from 8 to 6\nmove 6 from 4 to 1\nmove 14 from 3 to 9\nmove 1 from 6 to 4\nmove 3 from 3 to 9\nmove 1 from 4 to 5\nmove 10 from 9 to 6\nmove 6 from 6 to 7\nmove 2 from 1 to 8\nmove 1 from 8 to 6\nmove 16 from 3 to 2\nmove 1 from 8 to 1\nmove 1 from 7 to 1\nmove 7 from 3 to 4\nmove 1 from 6 to 5\nmove 4 from 2 to 3\nmove 5 from 4 to 9\nmove 2 from 4 to 5\nmove 4 from 7 to 4\nmove 5 from 9 to 6\nmove 2 from 5 to 4\nmove 11 from 6 to 7\nmove 1 from 6 to 8\nmove 5 from 1 to 5\nmove 2 from 6 to 4\nmove 7 from 7 to 3\nmove 1 from 8 to 6\nmove 2 from 7 to 3\nmove 1 from 1 to 3\nmove 3 from 2 to 8\nmove 9 from 2 to 5\nmove 1 from 6 to 1\nmove 1 from 4 to 8\nmove 7 from 4 to 7\nmove 8 from 5 to 6\nmove 1 from 7 to 2\nmove 1 from 7 to 4\nmove 3 from 7 to 8\nmove 1 from 2 to 3\nmove 1 from 1 to 2\nmove 1 from 1 to 7\nmove 3 from 7 to 6\nmove 11 from 6 to 2\nmove 4 from 8 to 7\nmove 2 from 8 to 7\nmove 15 from 3 to 2\nmove 7 from 9 to 4\nmove 3 from 3 to 2\nmove 4 from 4 to 7\nmove 5 from 7 to 3\nmove 3 from 4 to 6\nmove 3 from 6 to 9\nmove 1 from 4 to 2\nmove 1 from 8 to 1\nmove 2 from 3 to 7\nmove 2 from 3 to 7\nmove 23 from 2 to 5\nmove 1 from 9 to 1\nmove 1 from 7 to 9\nmove 1 from 1 to 8\nmove 8 from 7 to 1\nmove 1 from 8 to 4\nmove 1 from 4 to 2\nmove 3 from 9 to 8\nmove 1 from 7 to 9\nmove 22 from 5 to 9\nmove 1 from 8 to 5\nmove 1 from 7 to 4\nmove 1 from 4 to 5\nmove 1 from 8 to 3\nmove 2 from 9 to 3\nmove 5 from 5 to 2\nmove 5 from 5 to 4\nmove 3 from 2 to 7\nmove 1 from 7 to 3\nmove 6 from 1 to 7\nmove 4 from 3 to 1\nmove 6 from 2 to 8\nmove 1 from 5 to 6\nmove 2 from 8 to 1\nmove 12 from 9 to 4\nmove 8 from 9 to 4\nmove 1 from 2 to 9\nmove 2 from 9 to 8\nmove 3 from 2 to 8\nmove 5 from 8 to 6\nmove 7 from 7 to 1\nmove 4 from 8 to 9\nmove 1 from 6 to 1\nmove 17 from 4 to 7\nmove 1 from 2 to 4\nmove 2 from 4 to 1\nmove 6 from 4 to 6\nmove 1 from 1 to 4\nmove 7 from 1 to 5\nmove 9 from 7 to 9\nmove 8 from 9 to 8\nmove 5 from 8 to 3\nmove 1 from 5 to 6\nmove 2 from 3 to 6\nmove 1 from 9 to 1\nmove 1 from 6 to 1\nmove 10 from 6 to 1\nmove 1 from 5 to 1\nmove 2 from 9 to 1\nmove 1 from 9 to 7\nmove 2 from 6 to 8\nmove 2 from 8 to 2\nmove 1 from 6 to 8\nmove 22 from 1 to 9\nmove 9 from 7 to 5\nmove 1 from 8 to 1\nmove 2 from 8 to 3\nmove 4 from 5 to 9\nmove 1 from 8 to 3\nmove 5 from 1 to 9\nmove 2 from 7 to 3\nmove 2 from 4 to 7\nmove 1 from 8 to 5\nmove 2 from 2 to 4\nmove 1 from 5 to 8\nmove 9 from 5 to 8\nmove 2 from 7 to 5\nmove 2 from 4 to 5\nmove 3 from 8 to 4\nmove 3 from 4 to 3\nmove 2 from 8 to 6\nmove 1 from 6 to 4\nmove 3 from 5 to 9\nmove 1 from 6 to 3\nmove 12 from 3 to 5\nmove 1 from 3 to 1\nmove 7 from 5 to 4\nmove 1 from 1 to 3\nmove 1 from 8 to 1\nmove 7 from 5 to 1\nmove 6 from 9 to 6\nmove 29 from 9 to 5\nmove 2 from 4 to 6\nmove 26 from 5 to 2\nmove 24 from 2 to 7\nmove 1 from 3 to 2\nmove 8 from 1 to 7\nmove 7 from 6 to 9\nmove 2 from 5 to 3\nmove 1 from 6 to 4\nmove 3 from 8 to 5\nmove 2 from 3 to 8\nmove 2 from 2 to 8\nmove 5 from 9 to 2\nmove 27 from 7 to 2\nmove 2 from 8 to 3\nmove 2 from 9 to 5\nmove 3 from 8 to 5\nmove 2 from 7 to 4\nmove 3 from 4 to 7\nmove 2 from 3 to 2\nmove 4 from 5 to 1\nmove 5 from 7 to 2\nmove 29 from 2 to 8\nmove 9 from 8 to 3\nmove 2 from 4 to 8\nmove 7 from 3 to 2\nmove 3 from 5 to 4\nmove 1 from 7 to 5\nmove 3 from 5 to 6\nmove 2 from 1 to 8\nmove 2 from 6 to 8\nmove 3 from 4 to 2\nmove 4 from 4 to 2\nmove 1 from 6 to 8\nmove 8 from 2 to 4\nmove 2 from 3 to 5\nmove 1 from 4 to 1\nmove 3 from 1 to 2\nmove 4 from 8 to 2\nmove 3 from 4 to 9\nmove 3 from 4 to 1\nmove 2 from 9 to 5\nmove 1 from 4 to 6\nmove 4 from 5 to 1\nmove 1 from 6 to 8\nmove 1 from 9 to 3\nmove 4 from 2 to 3\nmove 15 from 8 to 2\nmove 9 from 8 to 1\nmove 1 from 3 to 9\nmove 5 from 1 to 9\nmove 3 from 9 to 7\nmove 2 from 7 to 6\nmove 3 from 3 to 2\nmove 1 from 7 to 8\nmove 1 from 9 to 6\nmove 1 from 9 to 8\nmove 2 from 8 to 2\nmove 1 from 1 to 2\nmove 1 from 3 to 7\nmove 4 from 1 to 7\nmove 19 from 2 to 5\nmove 1 from 1 to 4\nmove 1 from 7 to 4\nmove 1 from 1 to 5\nmove 3 from 1 to 4\nmove 1 from 1 to 8\nmove 6 from 2 to 4\nmove 7 from 2 to 1\nmove 2 from 7 to 9\nmove 8 from 2 to 8\nmove 2 from 7 to 3\nmove 1 from 6 to 4\nmove 10 from 4 to 6\nmove 5 from 6 to 7\nmove 2 from 9 to 8\nmove 6 from 8 to 9\nmove 1 from 2 to 3\nmove 2 from 8 to 3\nmove 5 from 1 to 8\nmove 8 from 5 to 2\nmove 8 from 8 to 7\nmove 7 from 2 to 8\nmove 1 from 1 to 2\nmove 1 from 9 to 7\nmove 1 from 4 to 2\nmove 2 from 2 to 6\nmove 5 from 9 to 3\nmove 2 from 8 to 6\nmove 2 from 3 to 9\nmove 4 from 8 to 6\nmove 7 from 6 to 1\nmove 8 from 1 to 5\nmove 1 from 8 to 7\nmove 1 from 9 to 6\nmove 12 from 5 to 3\nmove 1 from 4 to 8\nmove 2 from 9 to 5\nmove 1 from 2 to 3\nmove 3 from 5 to 1\nmove 1 from 1 to 5\nmove 21 from 3 to 8\nmove 2 from 1 to 5\nmove 6 from 5 to 7\nmove 2 from 5 to 6\nmove 10 from 6 to 9\nmove 1 from 6 to 8\nmove 13 from 8 to 2\nmove 2 from 5 to 4\nmove 2 from 4 to 3\nmove 4 from 9 to 1\nmove 5 from 7 to 8\nmove 12 from 8 to 1\nmove 5 from 9 to 6\nmove 1 from 3 to 7\nmove 2 from 6 to 5\nmove 11 from 2 to 1\nmove 1 from 8 to 4\nmove 16 from 1 to 9\nmove 1 from 2 to 6\nmove 1 from 8 to 5\nmove 12 from 9 to 3\nmove 14 from 7 to 2\nmove 1 from 7 to 9\nmove 1 from 4 to 2\nmove 1 from 7 to 5\nmove 3 from 9 to 5\nmove 4 from 6 to 9\nmove 3 from 9 to 4\nmove 1 from 8 to 4\nmove 2 from 4 to 5\nmove 1 from 7 to 1\nmove 5 from 3 to 5\nmove 2 from 4 to 2\nmove 8 from 2 to 7\nmove 7 from 2 to 4\nmove 1 from 3 to 7\nmove 3 from 9 to 7\nmove 2 from 2 to 9\nmove 3 from 4 to 5\nmove 6 from 1 to 8\nmove 6 from 1 to 5\nmove 3 from 9 to 2\nmove 22 from 5 to 9\nmove 1 from 5 to 6\nmove 2 from 2 to 3\nmove 5 from 7 to 6\nmove 5 from 8 to 9\nmove 2 from 7 to 2\nmove 20 from 9 to 4\nmove 1 from 8 to 3\nmove 2 from 2 to 5\nmove 1 from 2 to 5\nmove 15 from 4 to 8\nmove 1 from 5 to 7\nmove 6 from 9 to 1\nmove 5 from 4 to 8\nmove 2 from 4 to 8\nmove 1 from 2 to 1\nmove 5 from 6 to 5\nmove 5 from 5 to 7\nmove 1 from 9 to 8\nmove 5 from 7 to 2\nmove 2 from 5 to 1\nmove 4 from 7 to 5\nmove 1 from 5 to 9\nmove 1 from 6 to 8\nmove 1 from 7 to 2\nmove 6 from 3 to 4\nmove 3 from 5 to 7\nmove 1 from 9 to 2\nmove 6 from 2 to 3\nmove 1 from 3 to 4\nmove 13 from 8 to 9\nmove 7 from 1 to 5\nmove 6 from 9 to 2\nmove 1 from 1 to 4\nmove 6 from 2 to 3\nmove 1 from 1 to 4\nmove 5 from 9 to 7\nmove 11 from 8 to 4\nmove 7 from 7 to 3\nmove 2 from 7 to 8\nmove 1 from 8 to 2\nmove 8 from 4 to 1\nmove 2 from 1 to 6\nmove 2 from 5 to 8\nmove 3 from 1 to 9\nmove 1 from 8 to 2\nmove 11 from 3 to 2\nmove 2 from 8 to 9\nmove 9 from 4 to 7\nmove 11 from 3 to 8\nmove 7 from 9 to 6\nmove 5 from 4 to 6\nmove 3 from 7 to 3\nmove 1 from 7 to 1\nmove 5 from 7 to 6\nmove 2 from 3 to 5\nmove 1 from 3 to 4\nmove 5 from 2 to 5\nmove 1 from 1 to 7\nmove 1 from 4 to 8\nmove 1 from 7 to 6\nmove 7 from 5 to 7\nmove 2 from 5 to 7\nmove 3 from 1 to 7\nmove 1 from 2 to 3\nmove 1 from 6 to 4\nmove 1 from 3 to 4\nmove 1 from 5 to 3\nmove 18 from 6 to 4\nmove 9 from 7 to 1\nmove 14 from 4 to 6\nmove 3 from 6 to 4\nmove 12 from 6 to 7\nmove 2 from 5 to 3\nmove 3 from 7 to 4\nmove 6 from 4 to 7\nmove 5 from 1 to 7\nmove 5 from 4 to 5\nmove 5 from 2 to 1\nmove 9 from 8 to 4\nmove 9 from 1 to 3\nmove 2 from 8 to 2\nmove 4 from 2 to 4\nmove 1 from 7 to 6\nmove 1 from 2 to 3\nmove 1 from 8 to 9\nmove 1 from 6 to 9\nmove 2 from 9 to 3\nmove 3 from 4 to 1\nmove 13 from 3 to 5\nmove 12 from 5 to 1\nmove 7 from 1 to 8\nmove 1 from 3 to 6\nmove 4 from 5 to 4\nmove 1 from 5 to 2\nmove 8 from 4 to 9\n"
	utils.Run(solve, testInput, true)
	utils.Run(solve, realInput, true)
}
