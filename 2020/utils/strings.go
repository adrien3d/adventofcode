package utils

func StringArrayContains(elts []string, elt interface{}) bool {
	for _, element := range elts {
		if element == elt {
			return true
		}
	}
	return false
}