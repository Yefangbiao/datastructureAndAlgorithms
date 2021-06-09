package chapter17

import (
	"fmt"
	"testing"
)

func TestIncrement(t *testing.T) {
	A := new(Calc)
	Increment(A)
	fmt.Println(A)
	Increment(A)
	Increment(A)
	Increment(A)
	fmt.Println(A)
}
