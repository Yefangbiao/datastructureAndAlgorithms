package e5_16

import "testing"

func TestJoins(t *testing.T) {
	tests := []struct {
		strs []string
		sep  string
		want string
	}{
		{
			strs: []string{"a", "b", "c"},
			sep:  ",",
			want: "a,b,c",
		},
	}
	for _, test := range tests {
		if got := Joins(test.sep, test.strs...); got != test.want {
			t.Errorf("got = %v, want = %v", got, test.want)
		}
	}
}
