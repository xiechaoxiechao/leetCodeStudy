package Solutions

import "math"

func maxSubArray(nums []int) int {
	var dp = make([]int, len(nums))
	for i := range nums {
		dp[i] = nums[i]
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+dp[i], dp[i])
	}
	max := math.MinInt32
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
