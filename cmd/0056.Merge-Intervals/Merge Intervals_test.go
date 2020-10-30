package leetcode

import (
	"fmt"
	"testing"
)

type sample struct {
	paras
	ans
}

type paras struct {
	intervals [][]int
}

type ans struct {
	res [][]int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[][]int{{8, 10}, {2, 6}, {1, 3}, {15, 18}}},
			ans{[][]int{{1, 6}, {8, 10}, {15, 18}}},
		},
		{
			paras{[][]int{{1, 4}, {4, 5}}},
			ans{[][]int{{1, 5}}},
		},
		{
			// another test for quicksort
			paras{[][]int{{6, 4}, {1, 5}, {2, 2}, {7, 7}, {9, 9}, {3, 3}, {4, 4}, {5, 5}, {10, 10}, {8, 8}}},
			ans{[][]int{{1, 5}, {6, 4}, {7, 7}, {8, 8}, {9, 9}, {10, 10}}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, merge(p.intervals))
	}

}
