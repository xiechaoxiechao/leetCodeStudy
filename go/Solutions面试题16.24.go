package Solutions 

func pairSums(nums []int, target int) [][]int {
	sort.Ints(nums)
	var l, r = 0, len(nums) - 1
	var ans = make([][]int, 0, 32)
	for l < r {
		add := nums[l] + nums[r]
		if add == target {
			ans = append(ans, []int{nums[l], nums[r]})
			l++
			r--
		} else if add < target {
			l++
		} else {
			r--
		}
	}
	return ans

}
