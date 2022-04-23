package main

import "regexp"

func checkDNAInput(isiFile string) bool {
	var re = regexp.MustCompile(`^[AGCT]+$`)
	return re.MatchString(isiFile)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
