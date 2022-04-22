package main

import (
	"regexp"
)

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

//ACGT
func getL(str string) []int {
	var array []int = make([]int, 4)
	for i := 0; i < 4; i++ {
		array[i] = -1
	}

	for i := 0; i < len(str); i++ {
		if str[i] == 'A' {
			array[0] = i
		} else if str[i] == 'C' {
			array[1] = i
		} else if str[i] == 'G' {
			array[2] = i
		} else if str[i] == 'T' {
			array[3] = i
		}
	}
	return array
}

func changeToIndex(word string) int {
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

func boyerMoore(pattern, text string) int {
	var lastOcc []int = getL(pattern)
	var n int = len(text)
	var m int = len(pattern)
	var i int = m - 1
	var j int = m - 1
	for i <= n-1 {
		if pattern[j] == text[i] {
			if j == 0 {
				return i
			} else {
				i--
				j--
			}
		} else {
			var occ int = lastOcc[changeToIndex(text[i:i+1])]
			i = i + m - min(j, 1+occ)
			j = m - 1
		}
	}
	return -1
}

func main() {
	var text string = "AGCTAGCTCGTCGCTAGCTGTCGAGCTCGATCGCTGATAG"
	var pattern string = "TCGA"
	//fmt.Scanf("%s", &fileName)
	if checkDNAInput(text) {
		print(boyerMoore(pattern, text))
	} else {
		print("Input salah")
	}
}
