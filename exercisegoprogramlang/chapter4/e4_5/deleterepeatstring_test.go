package e4_5

import "testing"

func TestDeleteRepeatString(t *testing.T) {
	tests := []struct {
		arr  []string
		want []string
	}{
		{
			arr:  []string{"a", "v", "u", "l", "l", "c"},
			want: []string{"a", "v", "u", "l", "c"},
		},
	}
	for _, test := range tests {
		if got := DeleteRepeatString(test.arr); !compareSlice(got, test.want) {
			t.Errorf("errot, got = %v, want = %v", got, test.want)
		}
	}
}

func compareSlice(a1 []string, a2 []string) bool {
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
