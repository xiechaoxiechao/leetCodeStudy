/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package Solutions

// @lc code=start
func maxProfit_I(prices []int) int {
	max := -1
	ans := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] >= max {
			max = prices[i]
		} else {
			t := max - prices[i]
			if ans < t {
				ans = t
			}
		}
	}
	return ans
}

// @lc code=end
