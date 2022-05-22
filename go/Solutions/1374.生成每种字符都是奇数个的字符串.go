/*
 * @lc app=leetcode.cn id=1374 lang=golang
 *
 * [1374] 生成每种字符都是奇数个的字符串
 */
package Solutions

// @lc code=start
func generateTheString(n int) string {
	var res = make([]byte, n)
	for i := 0; i < 26 && i < n; i++ {
		res[i] = byte('a' + i)
	}
	if n <= 26 {
		return string(res[0:n])
	}
	var tem = n - 26
	if n%2 != 0 {
		res[25] = 'a'
	}
	for i := 0; i < tem; i++ {
		res[i+26] = 'a'
	}
	return string(res)
}

// @lc code=end
