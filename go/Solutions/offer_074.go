package Solutions


func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans = make([][]int, 0, 8)
	end := intervals[0][1]
	begin := intervals[0][0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			end = max_merge(intervals[i][1], end)
		} else {
			ans = append(ans, []int{begin, end})
			begin, end = intervals[i][0], intervals[i][1]
		}
	}
	ans = append(ans, []int{begin, end})
	return ans
}

func max_merge(i, j int) int {
	if i < j {
		return j
	}
	return i
}
