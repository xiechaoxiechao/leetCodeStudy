package Solutions


type Uf struct {
	master map[int]int
	rank   map[int]int
}

const NONE = -(1<<32 - 1)

func (u *Uf) find(i int) int {
	if v, ok := u.master[i]; ok {
		if v != i {
			u.master[i] = u.find(v)
		}
		return u.master[i]
	}
	return NONE

}

func (u *Uf) Union(i, j int) {
	x, y := u.find(i), u.find(j)
	if x == NONE || y == NONE {
		return
	}
	if x == y {
		return
	}
	if x > y {
		u.master[y] = x
		u.rank[x] = x - i + 1
	} else {
		u.master[x] = y
		u.rank[y] = y - i + 1
	}
}

func (u *Uf) isConnected(i, j int) bool {
	x, y := u.find(i), u.find(j)
	return x == y
}
func longestConsecutive(nums []int) int {
	var ans = -1
	var uf = &Uf{}
	uf.master = make(map[int]int, len(nums))
	uf.rank = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		uf.master[nums[i]] = nums[i]
		uf.rank[nums[i]] = 1
	}

	for i := 0; i < len(nums); i++ {
		uf.Union(nums[i], nums[i]+1)
	}
	for i := 0; i < len(nums); i++ {
		k := uf.find(nums[i])-nums[i]
		if k > ans {
			ans = k
		}

	}
	return ans+1
}
