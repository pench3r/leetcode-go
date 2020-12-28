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
	path string
}

type ans struct {
	target string
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{"/.hidden"},
			ans{"/.hidden"},
		},
		{
			paras{"/..hidden"},
			ans{"/..hidden"},
		},
		{
			paras{"/a/./b/../../c/"},
			ans{"/c"},
		},
		{
			paras{"/..."},
			ans{"/..."},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, simplifyPath(p.path))
	}

}
