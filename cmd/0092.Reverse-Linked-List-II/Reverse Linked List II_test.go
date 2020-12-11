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
	nums     []int
	one, two int
}

type ans struct {
	target []int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{1, 2, 3, 4, 5}, 2, 4},
			ans{[]int{1, 4, 3, 2, 5}},
		},
		{
			paras{[]int{1, 2, 3, 4, 5}, 2, 2},
			ans{[]int{1, 2, 3, 4, 5}},
		},
		{
			paras{[]int{1, 2, 3, 4, 5}, 1, 5},
			ans{[]int{5, 4, 3, 2, 1}},
		},
		{
			paras{[]int{1, 2, 3, 4, 5, 6}, 3, 4},
			ans{[]int{1, 2, 4, 3, 5, 6}},
		},
		{
			paras{[]int{3, 5}, 1, 2},
			ans{[]int{5, 3}},
		},
		{
			paras{[]int{3}, 3, 5},
			ans{[]int{3}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		head := structure.Int2List(p.nums)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, structure.List2Int(reverseBetween(head, p.one, p.two)))
	}

}
