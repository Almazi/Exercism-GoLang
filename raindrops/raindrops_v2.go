package raindrops

import "fmt"

const testVersion = 2

//Convert function is converting the factors of any given into four output
// string: Pling, Plang, Plong or the input in string format. If 3 is in the
// factors of the given number then output should show Pling, same goes for
// 5: Plang; 7: Plong.
func Convert(input int) string {
	output := string(input)
	if input%3 == 0 && input%5 == 0 && input%7 == 0 {
		output = "PlingPlangPlong"
		return output
	} else if input%3 == 0 && input%5 == 0 {
		output = "PlingPlang"
		return output
	} else if input%3 == 0 && input%7 == 0 {
		output = "PlingPlong"
		return output
	} else if input%5 == 0 && input%7 == 0 {
		output = "PlangPlong"
		return output
	} else if input%3 == 0 {
		output = "Pling"
		return output
	} else if input%5 == 0 {
		output = "Plang"
		return output
	} else if input%7 == 0 {
		output = "Plong"
		return output
	}
	output = fmt.Sprintf("%v", input)
	return output

}
