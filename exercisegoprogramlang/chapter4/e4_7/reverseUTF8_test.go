package e4_7

import "testing"

func TestReverseUTF8(t *testing.T) {
	tests := []struct {
		arr  string
		want string
	}{
		{
			arr:  "世界",
			want: "界世",
		},
	}
	for _, test := range tests {
		if got := ReverseUTF8([]byte(test.arr)); string(got) != test.want {
			t.Errorf("errot, got = %v, want = %v", got, test.want)
		}
	}
}
