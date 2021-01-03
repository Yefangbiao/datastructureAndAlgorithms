package e4_3

func Reverse(p *[10]int) {
	for i, j := 0, len(*p)-1; i <= j; i, j = i+1, j-1 {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}
