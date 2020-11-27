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
	one []int
	two int
}

type ans struct {
	ans int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{3, 1, 4, structure.NULL, 2}, 1},
			ans{1},
		},
		{
			paras{[]int{5, 3, 6, 2, 4, structure.NULL, structure.NULL, 1}, 3},
			ans{3},
		},
		{
			paras{[]int{41, 37, 44, 24, 39, 42, 48, 1, 35, 38, 40, structure.NULL, 43, 46, 49, 0, 2, 30, 36, structure.NULL, structure.NULL, structure.NULL, structure.NULL, structure.NULL, structure.NULL, 45, 47, structure.NULL, structure.NULL, structure.NULL, structure.NULL, structure.NULL, 4, 29, 32, structure.NULL, structure.NULL, structure.NULL, structure.NULL, structure.NULL, structure.NULL, 3, 9, 26, structure.NULL, 31, 34, structure.NULL, structure.NULL, 7, 11, 25, 27, structure.NULL, structure.NULL, 33, structure.NULL, 6, 8, 10, 16, structure.NULL, structure.NULL, structure.NULL, 28, structure.NULL, structure.NULL, 5, structure.NULL, structure.NULL, structure.NULL, structure.NULL, structure.NULL, 15, 19, structure.NULL, structure.NULL, structure.NULL, structure.NULL, 12, structure.NULL, 18, 20, structure.NULL, 13, 17, structure.NULL, structure.NULL, 22, structure.NULL, 14, structure.NULL, structure.NULL, 21, 23}, 25},
			ans{24},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		root := structure.Ints2TreeNode(p.one)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, kthSmallest(root, p.two))
	}

}
