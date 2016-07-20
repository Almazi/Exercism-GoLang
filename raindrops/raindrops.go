package raindrops

import "fmt"

const testVersion = 2

//Convert function is converting the factors of any given into four output
// string: Pling, Plang, Plong or the input in string format. If 3 is in the
// factors of the given number then output should show Pling, same goes for
// 5: Plang; 7: Plong.
//This version is the final one where i had to use only four conditions
func Convert(input int) string {
	output := ""
	if input%3 == 0 {
		output = "Pling"

	}
	if input%5 == 0 {
		output = output + "Plang"
	}
	if input%7 == 0 {
		output = output + "Plong"
		return output
	}
	if output == "" {
		output = fmt.Sprintf("%v", input)
	}
	return output

}
