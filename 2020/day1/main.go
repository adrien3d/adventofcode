package main

import (
	"fmt"
	"github.com/adrien3d/adventofcode/2020/utils"
)

func main() {
	testInput := "1721\n979\n366\n299\n675\n1456"
	//testRes := 1721*299

	testInParsed := utils.ParseIntList(testInput, "\n")
	for _, first := range testInParsed {
		for _, second := range testInParsed {
			if first+second == 2020 {
				fmt.Println(first, second)
			}
		}
	}
}
