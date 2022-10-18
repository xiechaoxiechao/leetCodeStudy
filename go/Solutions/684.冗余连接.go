/*
 * @lc app=leetcode.cn id=684 lang=golang
 *
 * [684] 冗余连接
 */
package Solutions

// @lc code=start
type uf struct {
	parent []int
	rank   []int
}

func (u uf) find(i int) int {
	if u.parent[i] != i {
		u.parent[i] = u.find(u.parent[i])
	}
	return u.parent[i]
}
func (u uf) union(i, j int) {
	x, y := u.find(i), u.find(j)
	if x == y {
		return
	}
	if u.rank[x] < u.rank[y] {
		u.parent[x] = y
	} else if u.rank[x] == u.rank[y] {
		u.parent[x] = y
		u.rank[y]++
	} else {
		u.parent[y] = x
	}
}

func NewUF(n int) *uf {
	u := &uf{}
	u.parent = make([]int, n)
	u.rank = make([]int, n)
	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.rank[i] = 1
	}
	return u
}
func (u uf) isConnected(i, j int) bool {
	return u.find(i) == u.find(j)
}
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	var uf = NewUF(n)
	for i := 0; i < n; i++ {
		x, y := edges[i][0]-1, edges[i][1]-1
		if uf.isConnected(x, y) {
			return edges[i]
		} else {
			uf.union(x, y)
		}
	}
	return []int{}
}

// @lc code=end
