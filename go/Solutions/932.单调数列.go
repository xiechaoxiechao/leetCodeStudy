/*
 * @lc app=leetcode.cn id=968 lang=golang
 *
 * [932] 漂亮数组
 */
package Solutions

// @lc code=start
func beautifulArray(n int) []int {

	var fil func(int) []int
	fil = func(i int) []int {
		a := make([]int, 0, i)
		if i == 1 {
			a = append(a, 1)
			return a
		}
		oddNum := (i + 1) / 2
		evenNum := i / 2
		for _, v := range fil(oddNum) {
			a = append(a, 2*v-1)
		}
		for _, v := range fil(evenNum) {
			a = append(a, 2*v)
		}
		return a

	}
	return fil(n)

}

// @lc code=end
