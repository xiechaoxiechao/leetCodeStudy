package Solutions

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax, res := 0, 0, 0

	for left < right {
		if leftMax < height[left] {
			leftMax = height[left]
		}
		if rightMax < height[right] {
			rightMax = height[right]
		}
		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}

	}
	return res
}

// @lc code=end
