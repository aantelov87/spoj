package main

import "fmt"

func decrypt(s string, col int) string {
	l := len(s)
	r := make([]byte, l)
	nRows := l / col
	c := 0
	for i := 0; i < col; i++ {
		for j := 0; j < nRows; j++ {
			cp := i
			if j%2 != 0 {
				cp = col - cp - 1
			}
			r[c] = s[j*col+cp]
			c++
		}
	}
	return string(r)
}

func main() {
	var n int
	var s string
	fmt.Scanf("%d", &n)
	for n != 0 {
		fmt.Scanf("%s", &s)
		fmt.Printf("%s\n", decrypt(s, n))
		fmt.Scanf("%d", &n)
	}
}
