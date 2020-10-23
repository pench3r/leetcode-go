package leetcode


func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	result, mp := []int{}, map[int]int{}
	for inx, v := range nums2 {
		mp[v] = inx
	}
	for _, v := range nums1 {
		inner_inx := mp[v]
		flag := false
		for i := inner_inx; i < len(nums2); i++ {
			if (nums2[i] > v) {
				flag = true
				result = append(result, nums2[i])
				break
			}
		}
		if flag == false {
			result = append(result, -1)
		}
	}
	return result
}