package leetcodes

import (
	"sort"
)


func MinimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	
	sort.Ints(nums)
	minValue := nums[k-1] - nums[0]

	for i := k; i < len(nums); i++ {
		if nums[i] - nums[i-k+1] < minValue {
			minValue = nums[i] - nums[i-k+1]
		}
	}
	return minValue
}
