/*
 * @lc app=leetcode.cn id=728 lang=golang
 *
 * [728] 自除数
 */

// @lc code=start
package Solutions

func selfDividingNumbers(left int, right int) []int {
	var ans = make([]int, 0, 100)
	for ; left <= right; left++ {
		t := left
		var j = 0
		for t > 0 {
			j = t % 10
			if j == 0 || left%j != 0 {
				break
			}
			t /= 10
		}
		if t == 0 {
			ans = append(ans, left)
		}
	}
	return ans
}

// @lc code=end
