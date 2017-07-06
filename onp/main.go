package main

import "fmt"

type stack struct {
	data  [500]byte
	count int
}

func (s *stack) top() byte {
	if s.count == 0 {
		return 0
	}
	return s.data[s.count-1]
}

func (s *stack) push(b byte) {
	s.data[s.count] = b
	s.count++
}

func (s *stack) pop() byte {
	if s.count == 0 {
		return 0
	}
	e := s.data[s.count-1]
	s.count--
	return e
}

func priority(op byte) int {
	switch op {
	case '-', '+':
		return 1
	case '/', '*':
		return 2
	case '^':
		return 3
	}
	return -1
}

func main() {
	var t int
	var in string
	var sm stack
	fmt.Scanf("%d", &t)
	for ; t > 0; t-- {
		fmt.Scanf("%s", &in)
		r := make([]byte, len(in))
		var rl int
		for i := range in {
			if in[i] >= 'a' && in[i] <= 'z' {
				r[rl] = in[i]
				rl++
			} else if in[i] == ')' {
				for c := sm.pop(); c != '(' && c != 0; c = sm.pop() {
					r[rl] = c
					rl++
				}
			} else {
				if in[i] == '(' {
					sm.push(in[i])
					continue
				}
				p1 := priority(in[i])
				for op := sm.top(); op != 0 && p1 < priority(op); op = sm.pop() {
					r[rl] = op
					rl++
				}
				sm.push(in[i])
			}
		}
		for op := sm.pop(); op != 0; op = sm.pop() {
			r[rl] = op
			rl++
		}
		fmt.Printf("%s\n", string(r[:rl]))
	}

}
