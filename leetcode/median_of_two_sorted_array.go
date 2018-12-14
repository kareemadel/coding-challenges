func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	i_min, i_max, half_len := 0, m, (m+n+1)/2
	var i, j int
	var median float64
	for i_min <= i_max {
		i = (i_min + i_max) / 2
		j = half_len - i
		if i > 0 && nums1[i-1] > nums2[j] {
			i_max = i - 1
		} else if i < m && nums2[j-1] > nums1[i] {
			i_min = i + 1
		} else {
			var max_of_left int
			if i == 0 {
				max_of_left = nums2[j-1]
			} else if j == 0 {
				max_of_left = nums1[i-1]
			} else {
				max_of_left = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				median = float64(max_of_left)
				break
			}
			var min_of_right int
			if i == m {
				min_of_right = nums2[j]
			} else if j == n {
				min_of_right = nums1[i]
			} else {
				min_of_right = min(nums1[i], nums2[j])
			}
			median = float64(max_of_left+min_of_right) / 2
			break
		}
	}
	return median
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
