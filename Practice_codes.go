package main

import (
	"fmt"
)

func fibo_loop(num int) {
	//p := fmt.Printf
	var first int
	var second int
	var next int

	for i := 1; i <= num; i++ {
		if i <= 1 {
			next = i
			second = i
		} else {
			next = first + second
			first = second
			second = next
		}
		//p("%d %d %d\n",first,second,next)

	}
}

func recur(n int) int {
	if n <= 1 {
		return n
	} else {

		return recur(n-1) + recur(n-2)
	}

}

func facto(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * facto(n-1)
	}
}
func main() {
	var n int
	fmt.Scanf("%d", &n)
	fibo_loop(n)
	fmt.Println(recur(n))
	fmt.Println(facto(n))
}
