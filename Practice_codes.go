package main

import (
	"fmt"
)

func main() {
	hour := 10
	minute := 0

	time := (hour*60 + minute) % (60*24)
	fmt.Printf("%02v:%02v \n", time, minute)

}

