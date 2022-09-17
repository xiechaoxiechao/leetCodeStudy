package Solutions 

func findErrorNums(nums []int) []int {
	var ans = [2]int{}
	for i := 0; i < len(nums); {
		if nums[i] != i+1 {
			if nums[i] == nums[nums[i]-1] {
				ans[0] = nums[i]
				i++
				ans[1] = i
				continue
			}
			t := nums[i]
			nums[i] = nums[t-1]
			nums[t-1] = t
		} else {
			i++
		}
	}
	return ans[:]

}
