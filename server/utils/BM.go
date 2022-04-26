package backend

//Mecari LastOccurence dari sebuah pattern
func GetL(pattern string) []int {
	var array []int = make([]int, 4)
	for i := 0; i < 4; i++ {
		array[i] = -1
	}

	for i := 0; i < len(pattern); i++ {
		if pattern[i] == 'A' {
			array[0] = i
		} else if pattern[i] == 'C' {
			array[1] = i
		} else if pattern[i] == 'G' {
			array[2] = i
		} else if pattern[i] == 'T' {
			array[3] = i
		}
	}
	return array
}

//String matching dengan Boyer-Moore
func BoyerMoore(pattern, text string) int {
	var lastOcc []int = GetL(pattern)
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
			var occ int = lastOcc[ChangeToIndex(text[i:i+1])]
			i = i + m - Min(j, 1+occ)
			j = m - 1
		}
	}
	return -1
}
