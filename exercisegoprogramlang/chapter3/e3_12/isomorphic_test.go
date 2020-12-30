package e3_12

import (
	"testing"
)

func TestIsomorphic(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want bool
	}{
		{"abc", "cba", true},
		{"你好世界", "世界你好", true},
		{"钟山风雨起苍黄", "百万雄师过大江", false},
		{"新的一年找到女朋友", "阿弥陀佛", false},
	}
	for _, test := range tests {
		if got := Isomorphic(test.s1, test.s2); got != test.want {
			t.Errorf("error, got = %v, want = %v", got, test.want)
		}
	}
}
