package utils

import (
	"strconv"
	"strings"
)

func ParseIntList(input, separator string) []int {
	lines := strings.Split(input, separator)
	list := make([]int, len(lines))
	for i, line := range lines {
		num, err := strconv.Atoi(line)
		CheckErr(err)
		list[i] = num
	}
	return list
}
