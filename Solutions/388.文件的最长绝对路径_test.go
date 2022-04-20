/*
 * @lc app=leetcode.cn id=388 lang=golang
 *
 * [388] 文件的最长绝对路径
 */

// @lc code=start
package Solutions

import "testing"

func Test_lengthLongestPath(t *testing.T) {
	lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext")
}
