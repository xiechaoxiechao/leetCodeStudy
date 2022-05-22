/*
 * @lc app=leetcode.cn id=2038 lang=golang
 *
 * [2038] 如果相邻两个颜色均相同则删除当前颜色
 */

// @lc code=start
package Solutions

func winnerOfGame(colors string) bool {
	var last byte = colors[0]
	var j = 0
	var aNum, bNum = 0, 0
	var i int
	for i = 1; i < len(colors); i++ {
		if colors[i] != last {
			if i-j > 2 {
				if last == 'A' {
					aNum += i - j - 2
				} else {
					bNum += i - j - 2
				}
			}

			last = colors[i]
			j = i
		}
	}
	if i != j && i-j > 2 {
		if last == 'A' {
			aNum += i - j - 2
		} else {
			bNum += i - j - 2
		}
	}
	if aNum > bNum {
		return true
	}
	return false
}

// @lc code=end
