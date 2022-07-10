/*
 * @lc app=leetcode.cn id=873 lang=golang
 *
 * [873] 最长的斐波那契子序列的长度
 */

// @lc code=start
package Solutions

func lenLongestFibSubseq(arr []int) int {
	length := len(arr)
	mp := make(map[int]int)
	ans := 0
	for i, v := range arr {
		mp[v] = i
	}
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}
	for i, v := range arr {
		for j := i - 1; j >= 0 && arr[j]*2 >
			v; j-- {
			if k, ok := mp[v-arr[j]]; ok {
				if dp[k][j]+1 > 3 {
					dp[j][i] = dp[k][j] + 1
				} else {
					dp[j][i] = 3
				}
				if dp[j][i] > ans {
					ans = dp[j][i]
				}
			}
		}
	}
	return ans

}

// func dfs(){

// }

// @lc code=end
//  2 3 4 5 6 10
