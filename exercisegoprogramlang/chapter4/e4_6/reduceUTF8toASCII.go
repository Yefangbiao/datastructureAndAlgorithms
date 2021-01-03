package e4_6

import "unicode"

func ReduceUTF8toASCII(s []rune) []rune {
	for i, r := range s {
		if unicode.IsSpace(r) {
			s[i] = ' '
		}
	}
	return s
}
