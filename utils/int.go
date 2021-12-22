package utils

import (
	"errors"
	"strconv"
)

func IntArrayContains(elts []int, elt interface{}) bool {
	for _, element := range elts {
		if element == elt {
			return true
		}
	}
	return false
}

func IntArrayMin(elts []int) (ret int) {
	ret = elts[0]
	for _, elt := range elts {
		if elt < ret {
			ret = elt
		}
	}
	return ret
}

func IntArrayMax(elts []int) (ret int) {
	for _, elt := range elts {
		if elt > ret {
			ret = elt
		}
	}
	return ret
}

func IntArrayRemoveDuplicates(intSlice []int) []int {
	allKeys := make(map[int]bool)
	var list []int
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func IntArrayRemove(s []int, index int) ([]int, error) {
	if index >= len(s) {
		return nil, errors.New("out of Range Error")
	}
	return append(s[:index], s[index+1:]...), nil
}

func BinaryIntFromString(s string) (r int) {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		Log(true, "err", "BinaryIntFromString error:", err)
	}
	return int(i)
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func IsLowest(data []int, compare int) bool {
	for _, elt := range data {
		if compare > elt {
			return false
		}
	}
	return true
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

func Sum(data []int) (ret int) {
	for _, elt := range data {
		ret += elt
	}
	return ret
}
