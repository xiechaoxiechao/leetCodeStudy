/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */
package Solutions

// @lc code=start
func maxProfit(prices []int) int {
	var dp = make([][3]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = [3]int{}
	}
	dp[0][2] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max_pri(dp[i-1][1], dp[i-1][0])
		dp[i][1] = dp[i-1][2] + prices[i]
		dp[i][2] = max_pri(dp[i-1][2], dp[i-1][0]-prices[i])
	}
	return max_pri(dp[len(dp)-1][1], dp[len(dp)-1][0])
}

func max_pri(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// @lc code=end
