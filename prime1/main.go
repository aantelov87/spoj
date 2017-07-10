package main

import (
	"fmt"
)

const max = 32226

type sieve struct {
	primes [max]int
	count  int
	max    int
}

func NewSieve(n int) *sieve {
	var s sieve
	s.max = n
	s.init()
	return &s
}

func (s *sieve) search(v int) int {
	min := 0
	max := s.count - 1
	var m = -1
	for min < max {
		m = (max + min) / 2
		if s.primes[m] == v {
			return m
		} else if s.primes[m] > v {
			max = m - 1
		} else {
			min = m + 1
		}
	}

	return m
}

func (s *sieve) init() {
	n := s.max + 1
	mark := make([]byte, n)
	mark[0] = 1
	mark[1] = 1
	for i := 3; i*i < n; i += 2 {
		if mark[i] != 0 {
			continue
		}
		for j := i * i; j < n; j += i {
			mark[j] = 1
		}
	}
	for i := range mark {
		if mark[i] == 0 && (i == 2 || i%2 != 0) {
			s.primes[s.count] = i
			s.count++
		}
	}
}

func (s *sieve) print(low, high int) {
	if low < s.max {
		idx := s.search(low)
		for i := idx; s.primes[i] <= high && s.primes[i] != 0; i++ {
			if s.primes[i] >= low {
				fmt.Printf("%d\n", s.primes[i])
			}
		}
		low = s.max + 1
	}
	if high <= s.max {
		return
	}
	l := high - low + 1
	mark := make([]byte, l)
	for i := 0; s.primes[i]*s.primes[i] <= high && i < s.count; i++ {
		r := (low / s.primes[i]) * s.primes[i]
		if r < low {
			r += s.primes[i]
		}
		for j := r; j <= high; j += s.primes[i] {
			mark[j-low] = 1
		}
	}
	for i := 0; i < l; i++ {
		if mark[i] == 0 {
			fmt.Printf("%d\n", i+low)
		}
	}

}

func main() {
	s := NewSieve(max)
	var n, l, h int
	fmt.Scanf("%d", &n)
	for ; n > 0; n-- {
		fmt.Scanf("%d%d", &l, &h)
		s.print(l, h)
		fmt.Printf("\n")
	}
}
