package leetcodes

func MaxProduct(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	dp := make([][2]int, length)
	if nums[0] < 0 {
		dp[0][1] = nums[0]
	} else {
		dp[0][0] = nums[0]
	}
	maxValue := nums[0]

	for j := 0; j < length; j++ {
		dp := make([][2]int, length)
		if nums[j] < 0 {
			dp[0][1] = nums[j]
		} else {
			dp[0][0] = nums[j]
		}

		if nums[j] > maxValue {
			maxValue = nums[j]
		}

		for i, k := j+1, 1; i < length; i, k = i+1, k+1 {
			if nums[i] <= 0 {
				if dp[k-1][1]*nums[i] >= dp[k-1][0] {
					dp[k][0] = dp[k-1][1] * nums[i]
				} else {
					dp[k][0] = 0
				}
				if maxValue < dp[k][0] {
					maxValue = dp[k][0]

				}
				dp[k][1] = dp[k-1][0] * nums[i]
			}
			if nums[i] > 0 {
				if dp[k-1][0]*nums[i] >= dp[k-1][0] {
					dp[k][0] = dp[k-1][0] * nums[i]
				} else {
					dp[k][0] = nums[i]
				}
				if maxValue < dp[k][0] {
					maxValue = dp[k][0]

				}
				dp[k][1] = dp[k-1][1] * nums[i]
			}
		}

	}

	return maxValue
}

func MaxProducts(nums []int) int {
	maxValue, minValue, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		mx, mn := maxValue, minValue
		maxValue = max(mx*nums[i], max(nums[i], nums[i]*mn))
		minValue = min(mn*nums[i], min(nums[i], nums[i]*mx))
		result = max(maxValue, result)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
