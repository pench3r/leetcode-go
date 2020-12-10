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
			paras{[]int{1, 2, 3, 4, 5}, 3},
			ans{[]int{3, 2, 1, 4, 5}},
		},
		{
			paras{[]int{1, 2, 3, 4, 5}, 2},
			ans{[]int{2, 1, 3, 4, 5}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		head := structure.Int2List(p.nums)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, structure.List2Int(reverseKGroup(head, p.two)))
	}

}
