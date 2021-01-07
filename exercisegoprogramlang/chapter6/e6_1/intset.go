// 函数6.1-6.4
package e6_1

import (
	"bytes"
	"fmt"
	"math"
)

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len 返回元素个数
func (s *IntSet) Len() int {
	count := 0
	for i := range (*s).words {
		for bit := 0; bit < 64; bit++ {
			if (*s).words[i]&(1<<bit) != 0 {
				count++
			}
		}
	}
	return count
}

// Remove 从集合去除元素x
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		return
	}
	s.words[word] &= (1 << bit) - 1
}

// Clear 清除所有元素
func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

// 返回集合的副本
func (s *IntSet) Copy() *IntSet {
	newSets := make([]uint64, len(s.words))
	copy(newSets, (*s).words)
	return &IntSet{newSets}
}

// AddAll 接收一串整型值作为参数
func (s *IntSet) AddAll(nums ...int) {
	if nums == nil {
		return
	}
	for _, num := range nums {
		s.Add(num)
	}
}

// OrWith s与t进行或操作
func (s *IntSet) OrWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			// 添加新的就只添加长度
			s.words = append(s.words, 0)
		}
	}
}

// IntersectionWith s与t进行交运算
func (s *IntSet) IntersectionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			// 添加新的就只添加长度
			s.words = append(s.words, 0)
		}
	}
}

// SubtractionWith s与t差运算
func (s *IntSet) SubtractionWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			for j := 0; j < 64; j++ {
				if s.words[word]&(1<<j) != 0 && t.words[word]&(1<<j) != 0 {
					s.words[word] &= math.MaxUint64 - (1 << j)
				}
			}
		} else {
			// 不用添加新长度
			break
		}
	}
}

// SymmetricaldifferenceWith s与t对称差运算
func (s *IntSet) SymmetricaldifferenceWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			for j := 0; j < 64; j++ {
				if s.words[word]&(1<<j) != 0 && t.words[word]&(1<<j) != 0 {
					s.words[word] &= math.MaxUint64 - (1 << j)
				} else {
					s.words[word] |= 1 << j
				}
			}
		} else {
			// 不用添加新长度
			break
		}
	}
}

// Elems 返回包含集合元素slice
func (s *IntSet) Elems() []uint64 {
	res := make([]uint64, s.Len())
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				res = append(res, uint64(i*64+j))
			}
		}
	}
	return res
}
