package leetcodes

import (
	"sort"
)


func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	
	sort.Ints(nums)
	minValue := nums[1] - nums[0]

	for i := 2; i < len(nums); i++ {
		if nums[i] - nums[i-1] < minValue {
			minValue = nums[i] - nums[i-1]
		}
	}
	return minValue
}