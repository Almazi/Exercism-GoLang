package main

import (
	"fmt"
)


func main() {

	p := fmt.Println
	input := 1
	output := string(input)
	if input % 3 == 0 {
		output = "Pling"
		p(output)
	} else if input % 5 == 0{
		output = "Plang"
		p(output)
	} else if input % 7 == 0{
		output = "Plong"
		p(output)
	}
	p(output)
}
