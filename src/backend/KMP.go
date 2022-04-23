package main

import "regexp"

func checkDNAInput(isiFile string) bool {
	var re = regexp.MustCompile(`^[AGCT]+$`)
	return re.MatchString(isiFile)
}

func border(pattern string) []int {
	var m int = len(pattern)
	var j int = 0
	var i int = 1
	var border []int = make([]int, m)
	border[0] = 0

	for i < m {
		if pattern[j] == pattern[i] {
			border[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = border[j-1]
		} else {
			border[i] = 0
			i++
		}
	}
	return border
}

func KMP(pattern, text string) int {
	var n int = len(text)
	var m int = len(pattern)
	var i int = 0
	var j int = 0
	var border []int = border(pattern)

	for i < n {
		if pattern[j] == text[i] {
			if j == m-1 {
				return i - m + 1
			}
			i++
			j++
		} else if j > 0 {
			j = border[j-1]
		} else {
			i++
		}
	}
	return -1
}

func main() {
	var text string = "AGCTAGCTCGTCGCTAGCTGTCGAGCTCGATCGCTGATAG"
	var pattern string = "TCGA"
	//fmt.Scanf("%s", &fileName)
	if checkDNAInput(text) {
		print(KMP(pattern, text))
	} else {
		print("Input salah")
	}
}
