package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(candidates, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	var combinationSumHelper func(candidates, currentCombination []int, sum int)
	combinationSumHelper = func(candidates, currentCombination []int, sum int) {
		if sum == target {
			validComb := make([]int, len(currentCombination))
			copy(validComb, currentCombination)
			result = append(result, validComb)
			return
		}
		for i, n := range candidates {
			if sum+n <= target {
				currentCombination = append(currentCombination, n)
				combinationSumHelper(candidates[i:], currentCombination, sum+n)
				currentCombination = currentCombination[:len(currentCombination)-1]
			} else {
				break
			}
		}
	}
	var buffer []int
	combinationSumHelper(candidates, buffer, 0)
	return result
}
