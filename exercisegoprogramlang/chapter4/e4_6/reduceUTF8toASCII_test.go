package e4_6

import "testing"

func TestReduceUTF8toASCII(t *testing.T) {
	tests := []struct {
		arr  []rune
		want []rune
	}{
		{
			arr:  []rune("hello,\nworld"),
			want: []rune("hello, world"),
		},
	}
	for _, test := range tests {
		if got := ReduceUTF8toASCII(test.arr); !compareSlice(got, test.want) {
			t.Errorf("errot, got = %v, want = %v", got, test.want)
		}
	}
}

func compareSlice(a1 []rune, a2 []rune) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i, v := range a1 {
		if a2[i] != v {
			return false
		}
	}
	return true
}
