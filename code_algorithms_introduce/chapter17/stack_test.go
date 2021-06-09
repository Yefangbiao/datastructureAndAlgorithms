package chapter17

import "testing"

func TestStack_MultiPop(t *testing.T) {
	s := &stack{s: []int{}}
	s.s = append(s.s, 1)
	s.s = append(s.s, 2)
	s.s = append(s.s, 3)
	s.s = append(s.s, 4)
	s.s = append(s.s, 5)
	s.MultiPop(6)
}
