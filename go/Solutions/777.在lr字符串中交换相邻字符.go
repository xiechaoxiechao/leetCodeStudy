/*
 * @lc app=leetcode.cn id=777 lang=golang
 *
 * [777] 在LR字符串中交换相邻字符
 */
package Solutions

import (
	"strings"
)

// @lc code=start
func canTransform(start string, end string) bool {
	n := len(start)
	startT := strings.ReplaceAll(start, "X", "")
	endT := strings.ReplaceAll(end, "X", "")
	if startT != endT {
		return false
	}
	lastI := 0
	j := 0
	i := 0
	for ind := 0; ind < len(startT); ind++ {

		for ; i < n; i++ {
			if start[i] == startT[ind] {
				break
			}
		}
		for ; j < n; j++ {
			if end[j] == startT[ind] {
				break
			}
		}
		if startT[ind] == 'R' {
			if i > j {
				return false
			} else {
				k := i
				for ; k < n; k++ {
					if start[k] == 'L' {
						break
					}
				}
				if k < j {
					return false
				}
			}
			lastI = j

		} else {
			if i < j {
				return false
			} else {
				if lastI > j {
					return false
				}
			}
			lastI = j
		}
		i++
		j++
	}
	return true

}

// @lc code=end
