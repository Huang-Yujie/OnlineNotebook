func sort(nums []int, l, r int) {
	if l >= r {
		return
	}
	k := partition(nums, l, r)
	sort(nums, l, k - 1)
	sort(nums, k + 1, r)
}

func partition(nums []int, l, r int) int {
	i, j := l, r
	pivot := nums[l]
	for {
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[j], nums[l] = nums[l], nums[j]
	return j
}