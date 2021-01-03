package e4_5

func DeleteRepeatString(strs []string) []string {
	res := strs[:0]
	for i := 0; i < len(strs); i++ {
		if i > 0 && strs[i] == strs[i-1] {
			continue
		}
		res = append(res, strs[i])
	}
	return res
}
