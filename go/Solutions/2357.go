package Solutions


func minimumOperations(nums []int) int {
	var mp = map[int]struct{}{}
	for _, v := range nums {
		if v != 0 {
			mp[v] = struct{}{}
		}
	}
	return len(mp)

}