package Solutions

func findMaxAverage(nums []int, k int) float64 {
	var sum = 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	var maxSum = sum
	for l, r := 0, k; r < len(nums); l, r = l+1, r+1 {
		sum += nums[r]
		sum -= nums[l]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)

}