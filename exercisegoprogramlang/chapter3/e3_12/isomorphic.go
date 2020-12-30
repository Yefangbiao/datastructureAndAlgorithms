package e3_12

func Isomorphic(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, s := range s1 {
		m1[s]++
	}
	for _, s := range s2 {
		m2[s]++
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	for k, v := range m2 {
		if m1[k] != v {
			return false
		}
	}

	return true
}
