package leetcode

import "sort"

func reorganizeString(S string) string {
	fs := frequencySort(S)
	if fs == "" {
		return ""
	}
	sb := []byte(fs)
	j := (len(sb)-1)/2 + 1
	ans := ""
	for i := 0; i <= (len(sb)-1)/2; i++ {
		ans += string(sb[i])
		if j < len(sb) {
			ans += string(sb[j])
		}
		j++
	}
	return ans
}

func frequencySort(S string) string {
	if S == "" {
		return ""
	}
	sMap := map[byte]int{}
	cMap := map[int][]byte{}
	sb := []byte(S)
	for _, s := range sb {
		sMap[s]++
		if sMap[s] > (len(sb)+1)/2 {
			return ""
		}
	}
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}
	var keys []int
	for key := range cMap {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	res := make([]byte, 0)
	for _, k := range keys {
		for i := 0; i < len(cMap[k]); i++ {
			for j := 0; j < k; j++ {
				res = append(res, cMap[k][i])
			}
		}
	}
	return string(res)
}
