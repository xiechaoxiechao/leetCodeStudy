/*
 * @lc app=leetcode.cn id=1487 lang=golang
 *
 * [1487] 保证文件名唯一
 */
package Solutions

import "strconv"

// @lc code=start
func getFolderNames(names []string) []string {
	var mp = make(map[string]int, len(names))
	var ans = make([]string, len(names))
	for i := 0; i < len(names); i++ {
		if count, ok := mp[names[i]]; ok {
			for {
				t := names[i] + "(" + strconv.Itoa(count) + ")"
				if _, ok := mp[t]; !ok {
					mp[names[i]] = count + 1
					mp[t] = 1
					ans[i] = t
					break
				}
				count++
			}
		} else {
			mp[names[i]] = 1
			ans[i] = names[i]
		}
	}
	return ans
}

// @lc code=end
