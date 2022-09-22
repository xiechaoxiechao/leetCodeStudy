package Solutions

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		if nums[i] == i+1 {
			i++
			continue
		}
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}

	}
	var ans = make([]int, 0, 8)
	for i, v := range nums {
		if i != v-1 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
