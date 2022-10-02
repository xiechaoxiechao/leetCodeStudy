/*
 * @lc app=leetcode.cn id=2001 lang=golang
 *
 * [2001] 可互换矩形的组数
 */
package Solutions

// @lc code=start
func interchangeableRectangles(rectangles [][]int) int64 {
	var mp = make(map[float64]int)
	for i := 0; i < len(rectangles); i++ {
		t := float64(rectangles[i][0]) / float64(rectangles[i][1])
		mp[t]++
	}
	var ans int64
	for _, v := range mp {
		v1 := int64(v)
		ans += (v1 - 1) * v1 / 2
	}
	return ans
}

// @lc code=end
