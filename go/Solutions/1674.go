package Solutions

func minMoves(nums []int, limit int) int {
	var dif = make([]int, limit*2+2)

	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r+1 {
		dif[2] += 2
		dif[1+min_minM(nums[l], nums[r])] -= 1
		dif[limit+max_minM(nums[l], nums[r])+1] += 1
		dif[nums[l]+nums[r]] -= 1
		dif[nums[l]+nums[r]+1] += 1
	}
	a := 0
	ans := 1<<31 - 1
	for i := 2; i < len(dif)-1; i++ {
		a += dif[i]
		if ans > a {
			ans = a
		}
	}
	return ans
}
func min_minM(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
func max_minM(i, j int) int {
	if i < j {
		return j
	}
	return i
}
