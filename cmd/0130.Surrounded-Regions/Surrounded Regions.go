package leetcode

import (
	"github.com/pench3r/leetcode-go/template"
)

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	rn, cn := len(board), len(board[0])
	uf := template.UnionFind{}
	uf.Init(rn*cn + 1)
	for r := 0; r < rn; r++ {
		for c := 0; c < cn; c++ {
			if (r == 0 || r == (rn-1) || c == 0 || c == (cn-1)) && board[r][c] == 'O' {
				uf.Union(r*cn+c, rn*cn)
			} else if board[r][c] == 'O' {
				if board[r-1][c] == 'O' {
					uf.Union((r-1)*cn+c, r*cn+c)
				}
				if board[r+1][c] == 'O' {
					uf.Union((r+1)*cn+c, r*cn+c)
				}
				if board[r][c-1] == 'O' {
					uf.Union(r*cn+c-1, r*cn+c)
				}
				if board[r][c+1] == 'O' {
					uf.Union(r*cn+c+1, r*cn+c)
				}
			}
		}
	}
	for r := 0; r < rn; r++ {
		for c := 0; c < cn; c++ {
			if uf.Find(r*cn+c) != uf.Find(rn*cn) {
				board[r][c] = 'X'
			}
		}
	}
}
