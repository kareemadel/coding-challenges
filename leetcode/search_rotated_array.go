func search(nums []int, target int) int {
	return searchHelper(nums, 0, len(nums)-1, target)
}

func searchHelper(nums []int, left, right, target int) int {
	if right < left {
		return -1
	}
	var mid int = (left + right) / 2
	if target == nums[mid] {
		return mid
	}
	if nums[left] <= nums[mid] {
		if target >= nums[left] && target < nums[mid] {
			return searchHelper(nums, left, mid-1, target)
		} else {
			return searchHelper(nums, mid+1, right, target)
		}
	} else {
		if target > nums[mid] && target <= nums[right] {
			return searchHelper(nums, mid+1, right, target)
		} else {
			return searchHelper(nums, left, mid-1, target)
		}
	}
	return -1
}