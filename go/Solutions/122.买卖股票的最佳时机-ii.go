/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */
package Solutions

// @lc code=start
func maxProfit_II(prices []int) int {
	var dp = make([][2]int, len(prices))
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max_mm(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(dp)-1][0]
}

func max_mm(i, j int) int {
	if i < j {
		return j
	}
	return i
}

// @lc code=end
