package e3_11

import (
	"log"
	"testing"
)

func TestComma(t *testing.T) {
	want := "12,345.124125"
	if got := Comma("12345.124125"); got != want {
		log.Fatalf("error, got = %v, want = %v", got, want)
	}
}
