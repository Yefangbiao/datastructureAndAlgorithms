package e4_7

func ReverseUTF8(b []byte) string {
	// rune: 32bit, byte: 8 bit

	r := []rune(string(b))
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
