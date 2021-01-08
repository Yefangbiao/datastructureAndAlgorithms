package uf

import "testing"

func TestUF(t *testing.T) {
	// 给定n个城市和nxn矩阵graph，如果graph[i][j]==1 代表两个城市连通
	tests := []struct {
		graph [][]int
		test  [][]int
		want  []bool
	}{
		{
			graph: [][]int{
				{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
			},
			test: [][]int{
				{0, 1}, {0, 2}, {1, 2},
			},
			want: []bool{true, false, false},
		},
	}
	for _, test := range tests {
		uf := Constructor(len(test.graph))
		for i := range test.graph {
			for j := range test.graph {
				if test.graph[i][j] == 1 {
					uf.Connect(i, j)
				}
			}
		}
		for i, v := range test.test {
			if got := uf.IsConnected(v[0], v[1]); got != test.want[i] {
				t.Errorf("uf_test: got = %t, want = %t", got, test.want[i])
			}
		}
	}
}
