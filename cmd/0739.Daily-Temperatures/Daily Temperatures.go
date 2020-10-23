package leetcode

func dailyTemperatures(T []int) []int {
	stack := make([]int, len(T))
	toCheck := []int{}
	for inx, v := range T {
		for len(toCheck) > 0 && v > T[toCheck[len(toCheck)-1]] {
			tmp := toCheck[len(toCheck)-1]
			stack[tmp] = inx - tmp
			toCheck = toCheck[:len(toCheck)-1]
		}
		toCheck = append(toCheck, inx)
	}
	return stack
}