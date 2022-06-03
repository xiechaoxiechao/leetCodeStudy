/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] 两个列表的最小索引总和
 */

// @lc code=start
package Solutions

func findRestaurant(list1 []string, list2 []string) []string {
	var mp = make(map[string]int)
	for i, v := range list1 {
		mp[v] = i
	}
	var min = 1 << 31
	var res = make([]string, 0)
	for i, v := range list2 {
		if value, ok := mp[v]; ok {
			add := value + i
			if add < min {
				min = add
				res = res[:0]
				res = append(res, v)
			} else if add == min {
				res = append(res, v)
			}
		}

	}
	return res
}

// @lc code=end
