/*
 * @lc app=leetcode.cn id=2039 lang=golang
 *
 * [2039] 网络空闲的时刻
 */

// @lc code=start
package Solutions

func networkBecomesIdle(edges [][]int, patience []int) int {
	var ed = make([][]int, len(patience))
	var status = make([]bool, len(patience))
	for i := 0; i < len(patience); i++ {
		ed[i] = make([]int, 0, 100)
	}
	for i := 0; i < len(edges); i++ {
		ed[edges[i][0]] = append(ed[edges[i][0]], edges[i][1])
		ed[edges[i][1]] = append(ed[edges[i][1]], edges[i][0])
	}

	var maxTime = 0
	var nod = make([]int, 0, 100)
	nod = ed[0]
	var depth = 1
	status[0] = true
	for len(nod) != 0 {
		var tem = make([]int, 0, 1025)
		for i := 0; i < len(nod); i++ {
			if status[nod[i]] {
				continue
			}
			status[nod[i]] = true
			time := depth * 2
			var t = time / patience[nod[i]]
			if time%patience[nod[i]] != 0 {
				t += 1
			}
			time = depth*2 + (t-1)*patience[nod[i]]
			if time > maxTime {
				maxTime = time
			}
			tem = append(tem, ed[nod[i]]...)
		}
		nod = tem
		depth++

	}

	return maxTime + 1
}

// @lc code=end
