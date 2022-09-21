package Solutions


func canPartitionKSubsets(nums []int, k int) bool {
	var sum = 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	avg := sum / k
	sort.Ints(nums)
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
	var buckets = make([]int, k)
	return backTrack(nums, 0, buckets, avg)

}

func backTrack(nums []int, index int, buckets []int, target int) bool {
	if index == len(nums) {
		return true
	}

	for i := 0; i < len(buckets); i++ {
		if i > 0 && buckets[i] == buckets[i-1] {
			continue
		}
		if nums[index]+buckets[i] <= target {
			buckets[i] += nums[index]
			if backTrack(nums, index+1, buckets, target) {
				return true
			}
			buckets[i] -= nums[index]
		}
	}
	return false
}
