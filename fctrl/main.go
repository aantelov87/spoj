package main

import "fmt"

func fctrlZeros(f int) (c int) {
	c = 0
	for f/5 != 0 {
		c += f / 5
		f /= 5
	}
	return
}

func main() {
	var t, v int
	fmt.Scanf("%d", &t)
	for ; t > 0; t-- {
		fmt.Scanf("%d", &v)
		fmt.Printf("%d\n", fctrlZeros(v))
	}
}
