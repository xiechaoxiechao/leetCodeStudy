/*
 * @lc app=leetcode.cn id=686 lang=golang
 *
 * [686] 重复叠加字符串匹配
 */
package Solutions

// @lc code=start
func repeatedStringMatch(a string, b string) int {
	var la = len(a)
	var lb = len(b)
	j := 0
	k := 0
	for i := 0; i < la; i++ {
		if a[i] == b[0] {
			j = i + 1
			k = 1
			for k < lb && a[j%la] == b[k] {
				j++
				k++
			}
			// fmt.Println(j)
			if k == lb {
				return (j-1)/la + 1
			}
		}
	}
	return -1

}

// @lc code=end
