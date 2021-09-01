package Solutions

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	len := len(nums)
	if len <= 3 {
		sum := 0
		for _, ele := range nums {
			sum += ele
		}
		return sum
	}
	var c, d, min int
	for i := 0; i < len; i++ {
		c = i
		min = nums[i]
		for j := i; j < len; j++ {
			if min > nums[j] {
				min = nums[j]
				c = j
			}
		}
		d = nums[i]
		nums[i] = nums[c]
		nums[c] = d
	}

	var j, k, add, abs, minadd, absab int
	minadd = nums[0] + nums[1] + nums[len-1]
	min = ab(minadd - target)
	if min == 0 {
		return minadd
	}
	for i := 0; i < len-1; i++ {
		for j, k = i+1, len-1; j != k; {
			add = nums[j] + nums[i] + nums[k]
			abs = add - target
			absab = ab(abs)
			if absab == 0 {
				return add
			}
			if absab < min {
				min = absab
				minadd = add
			}
			if abs >= 0 {
				k--
				continue
			}
			j++
			continue

		}

	}

	return minadd

}
func ab(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
