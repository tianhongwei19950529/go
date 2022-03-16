package main

/*
思路 会和后一位比较,如果比当前大 就需要减去当前的值,如果比当前小就需要加上.
*/

func romanToInt(s string) int {
	var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sum := 0

	for i := range s {
		value := symbolValues[s[i]]
		if i < len(s)-1 && value < symbolValues[s[i+1]] {
			sum -= value
		} else {
			sum += value
		}
	}
	return sum
}
func main() {

}
