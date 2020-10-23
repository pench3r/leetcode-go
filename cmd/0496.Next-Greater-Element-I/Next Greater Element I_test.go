package leetcode

import (
	"testing";
	"fmt";
)

type sample struct {
	paras
	ans
}

type paras struct {
	nums1 []int
	nums2 []int
}

type ans struct {
	num []int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{4, 1, 2}, []int{1, 3, 4, 2}},
			ans{[]int{-1, 3, -1}},
		},
		{
			paras{[]int{2, 4}, []int{1, 2, 3, 4}},
			ans{[]int{3, -1}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, nextGreaterElement(p.nums1, p.nums2))
	}

}