/*
 * @lc app=leetcode.cn id=2222 lang=golang
 *
 * [2222] 选择建筑的方案数
 */
package Solutions

// @lc code=start
func numberOfWays(s string) int64 {
	var countRight = make([]int, len(s))
	oneNum := 0
	zeroNum := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			zeroNum++
			countRight[i] = oneNum
		} else {
			oneNum++
			countRight[i] = zeroNum
		}
	}
	zeroNum = 0
	oneNum = 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zeroNum++
			ans += oneNum * countRight[i]
		} else {
			oneNum++
			ans += zeroNum * countRight[i]
		}
	}

	return int64(ans)

}

// @lc code=end
