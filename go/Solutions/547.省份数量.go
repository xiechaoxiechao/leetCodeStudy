/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */
package Solutions

// @lc code=start

// type uf struct {
// 	parent []int
// 	rank   []int
// }

// func (u uf) find(i int) int {
// 	if u.parent[i] != i {
// 		u.parent[i] = u.find(u.parent[i])
// 	}
// 	return u.parent[i]
// }
// func (u uf) union(i, j int) {
// 	x, y := u.find(i), u.find(j)
// 	if x == y {
// 		return
// 	}
// 	if u.rank[x] < u.rank[y] {
// 		u.parent[x] = y
// 	} else if u.rank[x] == u.rank[y] {
// 		u.parent[x] = y
// 		u.rank[y]++
// 	} else {
// 		u.parent[y] = x
// 	}
// }

//	func NewUF(n int) *uf {
//		u := &uf{}
//		u.parent = make([]int, n)
//		u.rank = make([]int, n)
//		for i := 0; i < n; i++ {
//			u.parent[i] = i
//			u.rank[i] = 1
//		}
//		return u
//	}
func findCircleNum(isConnected [][]int) int {
	var n = len(isConnected)
	var uf = NewUF(n)
	var ans = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	var pri = map[int]struct{}{}
	for i := 0; i < n; i++ {
		t := uf.find(i)
		if _, ok := pri[t]; !ok {
			ans++
			pri[t] = struct{}{}
		}
	}
	return ans
}

// @lc code=end
