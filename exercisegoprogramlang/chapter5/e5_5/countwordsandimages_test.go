package e5_5

import (
	"fmt"
	"testing"
)

func TestCountWordsAndImages(t *testing.T) {
	words, images, err := CountWordsAndImages("https://golang.org")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("words = %d, images = %d", words, images)
}
