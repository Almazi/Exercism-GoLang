package main

import (
	"fmt"
)

func main() {
	time := 24*60 + 0
	if time < 60*24 && time > 0 {
		fmt.Println("time is less than 1440")

	} else if time > 1439 {
		time = time % 1440
		fmt.Println("ok")

	} else if time < 0 {
		time = time%1440 + 60*24
		fmt.Println("time is less than 0")
	} else {

		fmt.Println(time * 0)
	}

}
