package main

import "fmt"

func main() {
	var v int
	fmt.Scanf("%d", &v)
	for v != 42 {
		fmt.Printf("%d\n", v)
		fmt.Scanf("%d", &v)
	}
}
