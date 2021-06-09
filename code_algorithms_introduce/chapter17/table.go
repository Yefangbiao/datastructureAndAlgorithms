package chapter17

type Table struct {
	size int   // 表大小
	num  int   // 表中元素个数
	t    []int // 表
}

// Insert 表满扩张
func (t *Table) Insert(x int) {
	if t.size == 0 {
		t.t = make([]int, 1)
		t.size = 1
	}

	if t.num == t.size {
		t.size *= 2
		newT := make([]int, t.size)
		copy(newT, t.t)
		t.t = newT
	}
	t.t[t.num] = x
	t.num++
}

// Delete 装载因子为四分之一的时候,收缩
func (t *Table) Delete() {
	t.num--
	if t.size == 0 {
		return
	}
	if t.num == 0 || t.size == 1 {
		t.size = 0
		t.num = 0
		t.t = nil
		return
	}
	if t.size%t.num == 0 && t.size/t.num == 4 {
		// 表收缩
		t.size = t.size / 4
		newTable := make([]int, t.size)
		for i := 0; i < t.size; i++ {
			newTable[i] = t.t[i]
		}
		t.t = newTable
	}
}
