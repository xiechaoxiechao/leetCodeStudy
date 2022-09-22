package Solutions


func subArrayRanges(nums []int) int64 {
	var ans = 0
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		max := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
                ans += (max - min)
                continue
			}
			if nums[j] > max {
				max = nums[j]
            
			}
             ans += (max - min)
			
		}
	}
	return int64(ans)
}