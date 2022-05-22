/*
 * @lc app=leetcode.cn id=388 lang=golang
 *
 * [388] 文件的最长绝对路径
 */

// @lc code=start
package Solutions

import (
	"bytes"
)

func lengthLongestPath(input string) int {
	var last = 0
	input = input + "\n"
	var stack = make([][]byte, 1024)
	var maxLen = 0
	stIndex := 0
	for i := 0; i < len(input); {

		if input[i] == '\t' {
			stIndex++
			i++
			last = i
		} else if input[i] == '\n' {

			st := []byte(input[last:i])
			stack[stIndex] = st
			stIndex++
			if bytes.IndexByte(st, '.') != -1 {
				var length = -1
				for j := 0; j < stIndex; j++ {
					length += len(stack[j]) + 1
				}
				if length > maxLen {
					maxLen = length
				}
			}
			stIndex = 0
			i++
			last = i

		} else {
			i++
		}
	}
	return maxLen

}

// @lc code=end
