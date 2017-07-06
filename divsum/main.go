package main

import (
	"fmt"
	"math"
)

func divSum(n int) (s int) {
	if n <= 1 {
		return 0
	}
	m := int(math.Sqrt(float64(n)))
	for i := 1; i <= m; i++ {
		if n%i == 0 {
			j := n / i
			if j != i && j != n {
				s += j
			}
			s += i
		}
	}
	return
}

func main() {
	var n, t int
	fmt.Scanf("%d", &t)
	for ; t > 0; t-- {
		fmt.Scanf("%d", &n)
		fmt.Printf("%d\n", divSum(n))
	}
}
