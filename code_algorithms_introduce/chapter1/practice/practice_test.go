package practice

import (
	"fmt"
	"testing"
)

func TestFindN_1_2_2(t *testing.T) {
	got := findN_1_2_2()
	want := 40
	if got != want {
		fmt.Printf("error,we want %d,we got %d\n", want, got)
	}
}

func TestFindN_1_2_3(t *testing.T) {
	got := findN_1_2_3()
	want := 15
	if got != want {
		fmt.Printf("error,we want %d,we got %d\n", want, got)
	}
}
