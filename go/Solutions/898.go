package Solutions

func subarrayBitwiseORs(arr []int) int {
	var mp1 = make(map[int]struct{})
	mp1[arr[0]] = struct{}{}
	var mp = make(map[int]struct{}, 32)
	mp[arr[0]] = struct{}{}
	for i := 1; i < len(arr); i++ {
		dpT := make(map[int]struct{})
		dpT[arr[i]] = struct{}{}
		mp[arr[i]] = struct{}{}
		for k := range mp1 {
			t := k | arr[i]
			mp[t] = struct{}{}
			dpT[t] = struct{}{}
		}
		mp1 = dpT

	}
	return len(mp)
}
