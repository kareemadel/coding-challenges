func twoSum(nums []int, target int) []int {
	intMap := make(map[int]int)
	var diff int
	for i := 0; i < len(nums); i++ {
		diff = target - nums[i]
		j, ok := intMap[diff]
		if ok {
			return []int{j, i}
		}
		intMap[nums[i]] = i
	}
	return []int{}
}