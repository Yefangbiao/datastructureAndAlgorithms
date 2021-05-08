// package uf is implement of union-find
package uf

// 路径压缩+按秩合并
type uf struct {
	id    []int // 分量id
	count int   // 分量个数
	sz    []int //各个根节点所对应的分量的大小
}

// 返回一个uf, n 为初始化的大小
func Constructor(n int) uf {
	newID := make([]int, n, n)
	newSZ := make([]int, n, n)
	for i := 0; i < n; i++ {
		newID[i] = i
		newSZ[i] = 1
	}
	return uf{id: newID, count: n, sz: newSZ}
}

// 返回分量个数
func (u *uf) Count() int { return u.count }

// 查询分量标识符
func (u *uf) Find(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

// 返回分量是否连通
func (u *uf) IsConnected(p, q int) bool { return u.Find(p) == u.Find(q) }

// 归并分量
func (u *uf) Connect(p, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)
	if pRoot == qRoot {
		return
	}
	if u.sz[pRoot] > u.sz[qRoot] {
		u.sz[pRoot] += u.sz[qRoot]
		u.id[qRoot] = pRoot
	} else {
		u.sz[qRoot] += u.sz[pRoot]
		u.id[pRoot] = qRoot
	}
	u.count--
}
