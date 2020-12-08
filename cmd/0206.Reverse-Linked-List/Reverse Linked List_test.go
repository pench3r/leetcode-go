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
}

type ans struct {
	target []int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{1, 2, 3, 4, 5}},
			ans{[]int{5, 4, 3, 2, 1}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		head := structure.Int2List(p.nums)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, structure.List2Int(reverseList(head)))
	}

}
