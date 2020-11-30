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
	nums []int
	two  int
}

type ans struct {
	target []int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{1, 4, 3, 2, 5, 2}, 3},
			ans{[]int{1, 2, 2, 4, 3, 5}},
		},
		{
			paras{[]int{1, 1, 2, 2, 3, 3, 3}, 2},
			ans{[]int{1, 1, 2, 2, 3, 3, 3}},
		},
		{
			paras{[]int{1, 4, 3, 2, 5, 2}, 0},
			ans{[]int{1, 4, 3, 2, 5, 2}},
		},
		{
			paras{[]int{4, 3, 2, 5, 2}, 3},
			ans{[]int{2, 2, 4, 3, 5}},
		},
		{
			paras{[]int{1, 1, 1, 1, 1, 1}, 1},
			ans{[]int{1, 1, 1, 1, 1, 1}},
		},
		{
			paras{[]int{3, 1}, 2},
			ans{[]int{1, 3}},
		},
		{
			paras{[]int{1, 2}, 3},
			ans{[]int{1, 2}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		head := structure.Int2List(p.nums)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, structure.List2Int(partition(head, p.two)))
	}

}
