package Solutions

import "math"

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	var flag = 1
	if dividend <= 0 {
		dividend = -dividend
		flag = -flag
	}
	if divisor <= 0 {
		divisor = -divisor
		flag = -flag
	}
	if dividend > math.MaxInt32 {
		if divisor == 1 && flag == -1 {
			return -math.MaxInt32 - 1
		} else if divisor == 1 && flag == 1 {
			return math.MaxInt32
		}

	}
	if dividend == math.MaxInt32 && divisor == 1 {
		if flag == 1 {
			return math.MaxInt32
		} else {
			return -math.MaxInt32
		}
	}
	var times, ri, res, l, r, mid int
	res = 0
	l = 0
	mid = 0
	for dividend >= divisor {
		times = 0
		if l == 0 {
			ri = divisor
			for ; ri <= dividend; times++ {
				ri <<= 1
			}
			l = times
			r = 0
			dividend -= ri >> 1
		} else {
			for {
				mid = (l-r)>>1 + r
				ri = divisor << mid
				if ri == dividend {
					times = mid + 1
					dividend = 0
					break
				} else if ri > dividend {
					l = mid
					continue
				} else {
					if (ri << 1) > dividend {
						dividend -= ri
						times = mid + 1
						l = times - 1
						r = 0
						break
					} else {
						r = mid + 1
					}
				}
			}
		}
		res += 1 << (times - 1)
	}
	if flag == 1 {
		return res
	} else {
		return -res
	}

}

// @lc code=end
