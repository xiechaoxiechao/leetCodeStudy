/*
 * @lc app=leetcode.cn id=1415 lang=golang
 *
 * [1415] 长度为 n 的开心字符串中字典序第 k 小的字符串
 */
package Solutions

// @lc code=start
func getHappyString(n int, k int) string {
	blockSize := (1 << (n - 1))
	if k > 3*blockSize {
		return ""
	}
	var ans = make([]byte, 0, n)
	for i := 0; i < 3; i++ {
		if k > i*blockSize && k <= (i+1)*blockSize {
			k -= i * blockSize
			ans = append(ans, byte('a'+i))
		}
	}
	blockSize /= 2
	for blockSize != 0 {
		c := ans[len(ans)-1]
		if k <= blockSize {

			if c == 'a' {
				ans = append(ans, 'b')
			} else if c == 'b' {
				ans = append(ans, 'a')
			} else {
				ans = append(ans, 'a')
			}

		} else {
			k -= blockSize
			if c == 'a' {
				ans = append(ans, 'c')
			} else if c == 'b' {
				ans = append(ans, 'c')
			} else {
				ans = append(ans, 'b')
			}
		}
		blockSize /= 2
	}
	return string(ans)
}

// @lc code=end
