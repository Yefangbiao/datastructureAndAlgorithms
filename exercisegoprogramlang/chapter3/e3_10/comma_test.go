package e3_10

import (
	"log"
	"testing"
)

func TestComma(t *testing.T) {
	want := "12,345"
	if got := Comma("12345"); got != want {
		log.Fatalf("error, got = %v, want = %v", got, want)
	}
}
