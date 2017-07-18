package main

import (
	"fmt"
	"math"
)

func main() {
	var t, n int
	fmt.Scanf("%d", &t)
	var r float64
	for ; t > 0; t-- {
		fmt.Scanf("%d", &n)
		if n <= 1 {
			fmt.Printf("1\n")
			continue
		}
		r = (math.Log(2*math.Pi) - 2*float64(n) + math.Log(float64(n))*(1+2*float64(n))) / (2 * math.Log(10))
		r = math.Ceil(r)
		fmt.Printf("%d\n", int(r))
	}
}
