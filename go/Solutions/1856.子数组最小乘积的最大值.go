/*
 * @lc app=leetcode.cn id=1856 lang=golang
 *
 * [1856] 子数组最小乘积的最大值
 */
package Solutions

// @lc code=start
func maxSumMinProduct(nums []int) int {
	length := len(nums)
	left := make([]int, length)
	right := make([]int, length)
	preSum := make([]int, length)
	var sum = 0
	for i, v := range nums {
		sum += v
		preSum[i] = sum
	}
	ms := make([]int, length)
	msTop := -1
	for i := 0; i < length; {
		if msTop == -1 {
			msTop = 0
			ms[0] = i
			i++
			continue
		}
		if nums[ms[msTop]] > nums[i] {
			right[ms[msTop]] = i - 1
			msTop--
		} else {
			msTop++
			ms[msTop] = i
			i++
		}
	}
	for msTop >= 0 {
		right[ms[msTop]] = length - 1
		msTop--
	}
	msTop = -1
	for i := length - 1; i >= 0; {
		if msTop == -1 {
			msTop = 0
			ms[0] = i
			i--
			continue
		}
		if nums[ms[msTop]] > nums[i] {
			left[ms[msTop]] = i + 1
			msTop--
		} else {
			msTop++
			ms[msTop] = i
			i--
		}
	}
	for msTop >= 0 {
		left[ms[msTop]] = 0
		msTop--
	}
	var ans int64 = 0

	for i, v := range nums {
		sum = 0
		if left[i] == 0 {
			sum = preSum[right[i]]
		} else {
			sum = preSum[right[i]] - preSum[left[i]-1]
		}
		t := (int64(sum) * int64(v))
		if ans < t {
			ans = t
		}
	}

	return int(ans % (1000000000 + 7))

}

// @lc code=end
