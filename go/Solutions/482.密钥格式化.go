/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 */
package Solutions

// @lc code=start
func licenseKeyFormatting(s string, k int) string {
	var l = len(s)
	var p = l / k
	l += p
	var res = make([]byte, l)
	l--
	n := 0
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == '-' {
			p++
			continue
		} else if s[i] > 96 {
			n++
			res[l] = s[i] - 32
			l--
		} else {
			n++
			res[l] = s[i]
			l--
		}
		if n == k {
			n = 0
			res[l] = '-'
			l--
			p--
		}
	}
	for i := 0; i < len(res); i++ {
		if res[i] > 47 {
			return string(res[i:])
		}
	}
	return ""

}

// @lc code=end
