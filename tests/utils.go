package tests

import "regexp"

func Split(s string) []string {
	re := regexp.MustCompile(`\s+|\n`)
	return re.Split(s, -1)
}

func IncludesSameMember(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	counts := make(map[string]int)

	for _, elem := range slice1 {
		counts[elem]++
	}

	for _, elem := range slice2 {
		counts[elem]--
		if counts[elem] < 0 {
			return false
		}
	}

	return true
}
