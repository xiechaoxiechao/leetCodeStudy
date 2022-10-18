/*
 * @lc app=leetcode.cn id=721 lang=golang
 *
 * [721] 账户合并
 */
package Solutions

import "sort"

// @lc code=start
type _uf struct {
	parent map[string]string
}

func (u _uf) find(s string) string {
	if v, ok := u.parent[s]; ok {
		if v != s {
			u.parent[s] = u.find(u.parent[s])
		}
		return u.parent[s]
	} else {
		u.parent[s] = s
		return s
	}
}
func (u _uf) union(s1, s2 string) {
	x, y := u.find(s1), u.find(s2)
	if x == y {
		return
	}
	u.parent[y] = x

}
func _newUf() *_uf {
	u := &_uf{}
	u.parent = make(map[string]string, 8)
	return u

}

func accountsMerge(accounts [][]string) [][]string {
	var uf = _newUf()
	name := make(map[string]string, len(accounts))
	for i := 0; i < len(accounts); i++ {
		uf.find(accounts[i][1])
		name[accounts[i][1]] = accounts[i][0]
		for j := 2; j < len(accounts[i]); j++ {
			uf.union(accounts[i][1], accounts[i][j])
		}
	}
	var ansMp = make(map[string][]string, 8)
	for k := range uf.parent {
		master := uf.find(k)
		ansMp[master] = append(ansMp[master], k)
	}
	var ans = make([][]string, 0, len(ansMp))
	for k, v := range ansMp {
		t := make([]string, 0, len(v)+1)
		t = append(t, name[k])
		sort.StringSlice(v).Sort()
		t = append(t, v...)
		ans = append(ans, t)
	}

	return ans
}

// @lc code=end
