package utils

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
