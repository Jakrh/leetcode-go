package solution

func merge(nums1, nums2 []int) []int {
	res := make([]int, len(nums1)+len(nums2))

	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			res[k] = nums1[i]
			k++
			i++
		} else {
			res[k] = nums2[j]
			k++
			j++
		}
	}

	for i < len(nums1) {
		res[k] = nums1[i]
		k++
		i++
	}

	for j < len(nums2) {
		res[k] = nums2[j]
		k++
		j++
	}

	return res
}

func sortArray(nums []int) []int {
	if len(nums) == 1 {
		return []int{nums[0]}
	}

	nums1 := nums[:len(nums)/2]
	nums2 := nums[len(nums)/2:]

	return merge(sortArray(nums1), sortArray(nums2))
}
