package chapter15

import "testing"

type steel struct {
	n    int
	want int
	name string
}

var tests = []struct {
	p      []int
	steels []steel
}{
	{
		p: []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30},
		steels: []steel{
			{n: 1, want: 1, name: "case1"},
			{n: 2, want: 5, name: "case2"},
			{n: 3, want: 8, name: "case3"},
			{n: 4, want: 10, name: "case4"},
			{n: 5, want: 13, name: "case5"},
			{n: 6, want: 17, name: "case6"},
			{n: 7, want: 18, name: "case7"},
			{n: 8, want: 22, name: "case8"},
			{n: 9, want: 25, name: "case9"},
			{n: 10, want: 30, name: "case10"},
		},
	},
}

func testFramework(t *testing.T, cutRod func([]int, int) int) {
	for _, test := range tests {
		for _, s := range test.steels {
			t.Run(s.name, func(t *testing.T) {
				if got := cutRod(test.p, s.n); got != s.want {
					t.Errorf("CutRod: got = %v, want = %v", got, s.want)
				}
			})
		}
	}
}

func TestCutRod(t *testing.T) {
	testFramework(t, CutRod)
}

func TestMemoizedCutRod(t *testing.T) {
	testFramework(t, MemoizedCutRod)
}

func TestBottomUpCutRod(t *testing.T) {
	testFramework(t, BottomUpCutRod)
}

func TestPrintCutRodSolution(t *testing.T) {
	for _, test := range tests {
		for _, s := range test.steels {
			t.Run(s.name, func(t *testing.T) {
				PrintCutRodSolution(test.p, s.n)
			})
		}
	}
}
