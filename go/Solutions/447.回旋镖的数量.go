/*
 * @lc app=leetcode.cn id=447 lang=golang
 *
 * [447] 回旋镖的数量
 */
package Solutions

// @lc code=start
func numberOfBoomerangs(points [][]int) int {
	var mps = make([]map[int][]int, len(points))
	for i := 0; i < len(points); i++ {
		mps[i] = make(map[int][]int, len(points)/2)
	}
	var ans = 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dis := dis(points[i], points[j])
			if v, ok := mps[i][dis]; ok {
				mps[i][dis] = append(v, j)
			} else {
				mps[i][dis] = []int{j}
			}
			if v, ok := mps[j][dis]; ok {
				mps[j][dis] = append(v, i)
			} else {
				mps[j][dis] = []int{i}
			}
		}
		for _, v := range mps[i] {
			ans += len(v) * (len(v) - 1)
		}
	}
	return ans
}
func dis(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}

// @lc code=end
// a 0
// a b 2
// a b c 6
// a b c d 12
