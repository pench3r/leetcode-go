package leetcode

import (
	"strconv"
)

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	numsStr := arrayToString(nums)
	quickSortString(numsStr, 0, len(numsStr)-1)
	res := ""
	for _, nStr := range numsStr {
		if nStr == "0" && res == "0" {
			continue
		}
		res = res + nStr
	}
	return res
}

func arrayToString(nums []int) []string {
	res := make([]string, 0)
	for _, num := range nums {
		res = append(res, strconv.Itoa(num))
	}
	return res
}

func quickSortString(numsStr []string, start int, end int) {
	if start >= end {
		return
	}
	p := parttionString(numsStr, start, end)
	quickSortString(numsStr, start, p-1)
	quickSortString(numsStr, p+1, end)
}

func parttionString(numsStr []string, start int, end int) int {
	left := start
	right := end
	piovt := end
	for left < right {
		for (numsStr[piovt]+numsStr[left]) <= (numsStr[left]+numsStr[piovt]) && left < right {
			left++
		}
		for (numsStr[piovt]+numsStr[right]) >= (numsStr[right]+numsStr[piovt]) && left < right {
			right--
		}
		if left < right {
			numsStr[left], numsStr[right] = numsStr[right], numsStr[left]
		}
	}
	numsStr[left], numsStr[piovt] = numsStr[piovt], numsStr[left]
	return left
}
