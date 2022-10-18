/*
 * @lc app=leetcode.cn id=821 lang=golang
 *
 * [821] 字符的最短距离
 */
package Solutions

// @lc code=start
const MAX = 1<<32 - 1

func shortestToChar(s string, c byte) []int {
	var ans = make([]int, len(s))
	if s[0] == c {
		ans[0] = 0
	} else {
		ans[0] = MAX
	}
	for i := 1; i < len(s); i++ {
		if s[i] == c {
			ans[i] = 0
		} else {
			if ans[i-1] == MAX {
				ans[i] = MAX
			} else {
				ans[i] = ans[i-1] + 1
			}
		}
	}
	last := -1
	if s[len(s)-1] == c {
		last = 0
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == c {
			last = 0
		} else {
			if last != -1 {
				last++
				if last < ans[i] {
					ans[i] = last
				}
			}

		}
	}
	return ans
}

// @lc code=end
