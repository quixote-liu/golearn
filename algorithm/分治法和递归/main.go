package main

import "fmt"

func main() {
	// fmt.Println(Rescuvie(10))
	fmt.Println(RescuvieTail(5, 5))
}

func Rescuvie(n int) int {
	if n == 0 {
		return 1
	}
	return n + Rescuvie(n-1)
}

func RescuvieTail(n int, a int) int {
	if n == 0 {
		return a
	}

	return RescuvieTail(n-1, a*n)
}
