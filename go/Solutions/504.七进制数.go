/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 */

// @lc code=start
package Solutions

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var ans = [9]byte{}
	var flag bool
	if num < 0 {
		flag = true
		num *= -1
	}
	tem := 7 * 7 * 7 * 7 * 7 * 7 * 7 * 7
	var offset = -1
	for i := 0; i < 9; i++ {
		if tem < num {
			k := num / tem
			ans[i] = byte(k + 48)
			num -= k * tem

			if offset == -1 {
				offset = i
			}
		}
		tem /= 7
	}
	if flag {
		return "-" + string(ans[offset:])
	}
	return string(ans[offset:])
}

// @lc code=end
