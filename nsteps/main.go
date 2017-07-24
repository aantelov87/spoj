package main

import "fmt"

func steps(x, y int) int {
	d := x - y
	if d == 0 {
		return 2*x - x%2
	}
	if d == 2 {
		return x + y - x%2
	}
	return -1
}

func main() {
	var n, x, y int
	fmt.Scanf("%d", &n)
	for n > 0 {
		fmt.Scanf("%d%d", &x, &y)
		r := steps(x, y)
		if r >= 0 {
			fmt.Printf("%d\n", r)
		} else {
			fmt.Println("No Number")
		}
		n--
	}
}
