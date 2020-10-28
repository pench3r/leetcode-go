package leetcode

import (
	"github.com/pench3r/leetcode-go/template"
)

func findCircleNum(M [][]int) int {
	uf := template.UnionFind{}
	uf.Init(len(M))
	for i := 0; i < len(M); i++ {
		for j := 0; j <= i; j++ {
			if M[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.TotalCount()
}
