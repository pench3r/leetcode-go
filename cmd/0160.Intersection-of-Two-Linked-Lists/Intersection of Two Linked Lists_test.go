package leetcode

import (
	"fmt"
	"testing"

	"github.com/pench3r/leetcode-go/structure"
)

type sample struct {
	paras
	ans
}

type paras struct {
	l1 []int
	l2 []int
}

type ans struct {
	res []int
}

// problem: 环链表的实现
// 测试无法通过，因为这里实现的两个单链表并非真正的交叉
// 直接在leetcode上测试即可
func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{}, []int{}},
			ans{[]int{}},
		},
		{
			paras{[]int{3}, []int{1, 2, 3}},
			ans{[]int{3}},
		},
		{
			paras{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans{[]int{1, 2, 3, 4}},
		},
		{
			paras{[]int{4, 1, 8, 4, 5}, []int{5, 0, 1, 8, 4, 5}},
			ans{[]int{8, 4, 5}},
		},
		{
			paras{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans{[]int{}},
		},
		{
			paras{[]int{0, 9, 1, 2, 4}, []int{3, 2, 4}},
			ans{[]int{2, 4}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		l1 := structure.Int2List(p.l1)
		l2 := structure.Int2List(p.l2)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, structure.List2Int(getIntersectionNode(l1, l2)))
	}
}
