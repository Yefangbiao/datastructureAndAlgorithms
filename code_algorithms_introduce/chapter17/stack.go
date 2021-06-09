package chapter17

import "log"

type stack struct {
	s []int
}

func (s *stack) Push(x int) {
	s.s = append(s.s, x)
}

func (s *stack) Pop() int {
	x := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return x
}

func (s *stack) MultiPop(k int) {
	for len(s.s) > 0 && k > 0 {
		x := s.Pop()
		log.Println("pop", x)
		k--
	}
}
