package e5_9

import "testing"

func TestExpand(t *testing.T) {
	tests := []struct {
		test string
		want string
	}{
		{
			test: "hello,world!",
			want: "hello,world!",
		},
		{
			test: "$fooaskbhuobehanh$fooasfbhj",
			want: "abcaskbhuobehanhabcasfbhj",
		},
	}
	for _, test := range tests {
		if got := Expand(test.test, replace); got != test.want {
			t.Errorf("error, got = %v, want = %v", got, test.want)
		}
	}
}
