package main

import "fmt"

func reverse(n int) int {
	var r int = 0
	var d int
	for n > 0 {
		d = n % 10
		r *= 10
		r += d
		n /= 10
	}
	return r
}

func main() {
	var n, a, b int
	fmt.Scanf("%d", &n)
	for n > 0 {
		fmt.Scanf("%d%d", &a, &b)
		v := reverse(a) + reverse(b)
		fmt.Printf("%d\n", reverse(v))
		n--
	}

}
