package backend

import "regexp"

// Sanitasi DNA dengan regex
func CheckDNAInput(isiFile string) bool {
	var re = regexp.MustCompile(`^[AGCT]+$`)
	return re.MatchString(isiFile)
}

// Cari nilai minimal dari 2 integer
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//Mengubah huruf A, C, G, dan T menjadi index
func ChangeToIndex(word string) int {
	if word == "A" {
		return 0
	} else if word == "C" {
		return 1
	} else if word == "G" {
		return 2
	} else if word == "T" {
		return 3
	}
	return -1
}
