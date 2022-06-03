/*
 * @lc app=leetcode.cn id=1928 lang=golang
 *
 * [1928] 规定时间内到达终点的最小花费
 */

// @lc code=start
package Solutions

import "math"

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	var ed = make([][]int, len(passingFees))
	var timp = make([][]int, len(passingFees))
	for i := 0; i < len(passingFees); i++ {
		timp[i] = make([]int, len(passingFees))
	}
	for _, v := range edges {
		ed[v[0]] = append(ed[v[0]], v[1])
		ed[v[1]] = append(ed[v[1]], v[0])
		timp[v[0]][v[1]] = v[2]
		timp[v[1]][v[0]] = v[2]
	}
	var cos = make([]int, len(passingFees))
	for i := 0; i < len(cos); i++ {
		cos[i] = math.MaxInt32
	}
	var time = make([]int, len(passingFees))
	time[0] = 0
	cos[0] = passingFees[0]
	var mp = make(map[int]struct{})
	mp[0] = struct{}{}
	for {
		var newmp = make(map[int]struct{})
		for i := range mp {

			for _, v := range ed[i] {
				if cos[v] >= cos[i]+passingFees[i] {
					if a := time[i] + timp[i][v]; a <= maxTime {
						cos[v] = cos[i] + passingFees[v]
						time[v] = time[i] + timp[i][v]
						newmp[v] = struct{}{}
					}
				}
			}
		}
		if len(newmp) == 0 {
			break
		}
		mp = newmp

	}
	if cos[len(passingFees)-1] == math.MaxInt32 {
		return -1
	}
	return cos[len(passingFees)-1]
}

// @lc code=end
