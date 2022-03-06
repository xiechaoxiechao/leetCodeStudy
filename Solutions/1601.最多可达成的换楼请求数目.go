/*
 * @lc app=leetcode.cn id=1601 lang=golang
 *
 * [1601] 最多可达成的换楼请求数目
 */

// @lc code=start
package Solutions

func maximumRequests(n int, requests [][]int) int {
	var status = make([]int, n)
	var dfs func(int, int)
	var maxReqNum = 0
	dfs = func(begin int, l int) {
		for i := begin; i < len(requests); i++ {
			//xuanzhong
			l++
			status[requests[i][0]]--
			status[requests[i][1]]++
			flag := 0
			for _, v := range status {
				flag |= v
			}
			if flag == 0 {

				if l > maxReqNum {
					maxReqNum = l
				}

			}
			dfs(i+1, l)
			status[requests[i][0]]++
			status[requests[i][1]]--
			l--
			//buxuanzhong
			dfs(i+1, l)
		}
	}
	dfs(0, 0)
	return maxReqNum

}

// @lc code=end
