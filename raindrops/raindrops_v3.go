package raindrops

import "strconv"

const testVersion = 2

//Convert function is converting the factors of any given into four output
// string: Pling, Plang, Plong or the input in string format. If 3 is in the
// factors of the given number then output should show Pling, same goes for
// 5: Plang; 7: Plong.
func Convert(num int) string {
	translation := pling(num) + plang(num) + plong(num)
	if len(translation) == 0 {
		translation = strconv.Itoa(num)
	}
	return translation
}

// pling returns correct string value for number divisible by 3
func pling(num int) string {
	if num%3 == 0 {
		return "Pling"
	}
	return ""
}

// plang returns correct string value for number divisible by 5
func plang(num int) string {
	if num%5 == 0 {
		return "Plang"
	}
	return ""
}

// plong returns correct string value for number divisible by 7
func plong(num int) string {
	if num%7 == 0 {
		return "Plong"
	}
	return ""
}