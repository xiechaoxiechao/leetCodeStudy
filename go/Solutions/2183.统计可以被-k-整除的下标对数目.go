/*
 * @lc app=leetcode.cn id=2183 lang=golang
 *
 * [2183] 统计可以被 K 整除的下标对数目
 */

// @lc code=start
package Solutions

func countPairs(nums []int, k int) int64 {
	return 2
}
func getMaxFactor(i, k int) int {
	var ans = 1
	var t int
	if i < k {
		t = i
	} else {
		t = k
	}
	for j := 2; j <= t; {
		if i/j == 0 && k/j == 0 {
			ans *= j
			i /= j
			k /= j
			t /= j
		} else {
			j += 1
		}
	}
	return ans

}

// @lc code=end
