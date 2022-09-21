package Solutions

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}
   
	sort.Ints(nums)
	var buckets = make([][]int, 1)
	buckets[0] = make([]int, 0, k)
	buckets[0] = append(buckets[0], nums[0])
	for i := 1; i < len(nums); i++ {
		j := 0
		for ; j < len(buckets); j++ {
			if len(buckets[j]) < k && buckets[j][len(buckets[j])-1]+1 == nums[i] {
				buckets[j] = append(buckets[j], nums[i])
				break
			}
		}
		if j == len(buckets) {
			buckets = append(buckets, make([]int, 0, k))
			buckets[len(buckets)-1] = append(buckets[len(buckets)-1], nums[i])
		}

	}
	for i := 0; i < len(buckets); i++ {
		if len(buckets[i]) != k {
			return false
		}
	}
	return true

}
