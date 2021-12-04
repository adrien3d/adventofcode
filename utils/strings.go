package utils

import (
	"errors"
)

func StringArrayContains(elts []string, elt interface{}) bool {
	for _, element := range elts {
		if element == elt {
			return true
		}
	}
	return false
}

func Remove(s []string, index int) ([]string, error) {
	if index >= len(s) {
		return nil, errors.New("out of Range Error")
	}
	return append(s[:index], s[index+1:]...), nil
}
