/*
 * @lc app=leetcode.cn id=1869 lang=golang
 *
 * [1869] 哪种连续子字符串更长
 */

// @lc code=start
package Solutions

func checkZeroOnes(s string) bool {
	var l1, l2 int
	var last byte = '0'
	var l = 0
	for i := 0; i < len(s); i++ {
		if s[i] == last {
			l++
		} else {
			if last == '0' {
				last = '1'
				if l > l1 {
					l1 = l
				}
			} else {
				last = '0'
				if l > l2 {
					l2 = l
				}
			}
			l = 1
		}

	}
	if last == '0' {
		last = '1'
		if l > l1 {
			l1 = l
		}
	} else {
		last = '0'
		if l > l2 {
			l2 = l
		}
	}
	return l2 > l1
}

// @lc code=end
