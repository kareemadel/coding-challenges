func subsets(nums []int) [][]int {
	var powerSet [][]int
	powerSet = append(powerSet, make([]int, 0))
	var subsetsHelper func(nums, subset []int)
	subsetsHelper = func(nums, subset []int) {
		if len(nums) == 0 {
			return
		}
		for i, v := range nums {
			newSubset := make([]int, len(subset)+1)
			tempSubset := append(subset, v)
			copy(newSubset, tempSubset)
			powerSet = append(powerSet, newSubset)
			subsetsHelper(nums[i+1:], tempSubset)
		}
		return
	}
	var buffer []int
	subsetsHelper(nums, buffer)
	return powerSet
}