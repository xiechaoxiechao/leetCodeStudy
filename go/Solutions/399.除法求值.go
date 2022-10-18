/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */
package Solutions

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var mp = make(map[string]map[string]float64, len(equations))
	for i, v := range equations {
		if mp[v[0]] == nil {
			mp[v[0]] = map[string]float64{
				v[1]: values[i],
			}
		} else {
			mp[v[0]][v[1]] = values[i]
		}
		if mp[v[1]] == nil {
			mp[v[1]] = map[string]float64{
				v[0]: 1 / values[i],
			}
		} else {
			mp[v[1]][v[0]] = 1 / values[i]
		}

	}
	var dfs func(from string, to string, last map[string]struct{}) float64

	dfs = func(from, to string, last map[string]struct{}) float64 {
		if v, ok := mp[from][to]; ok {
			return v
		}
		last[from] = struct{}{}
		for k, v := range mp[from] {
			if _, ok := last[k]; ok {
				continue
			}
			n := dfs(k, to, last)
			if n != -1 {
				return v * n
			}
		}
		delete(last, from)
		return -1

	}
	var ans = make([]float64, 0, len(queries))
	for _, v := range queries {
		ans = append(ans, dfs(v[0], v[1], map[string]struct{}{}))
	}
	return ans
}

// @lc code=end
