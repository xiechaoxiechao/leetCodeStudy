/*
 * @lc app=leetcode.cn id=1464 lang=golang
 *
 * [1464] 数组中两元素的最大乘积
 */
package Solutions

// @lc code=start
func maxProduct(nums []int) int {
	return (findMax(nums) - 1) * (findMax(nums) - 1)
}

func findMax(arr []int) int {
	ind := 0
	maxValue := arr[0]
	for i := 1; i < len(arr); i++ {
		if maxValue < arr[i] {
			maxValue = arr[i]
			ind = i
		}
	}
	arr[ind] = -1
	return maxValue
}

// @lc code=end
