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
	board [][]byte
}

type ans struct {
	board [][]byte
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[][]byte{}},
			ans{[][]byte{}},
		},
		{
			paras{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}},
			ans{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		solve(p.board)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, p.board)
	}

}