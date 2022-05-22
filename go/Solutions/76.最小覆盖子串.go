package Solutions

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	var MapWin = make([]int, 128)
	var MapAim = make([]int, 128)
	for i := 0; i < len(t); i++ {
		MapWin[t[i]] = 0
		MapAim[t[i]] += 1
	}
	var left = 0
	var right = 0
	var compare = func(a, b []int, t string) bool {
		for i := 0; i < len(t); i++ {
			if a[t[i]] > b[t[i]] {
				return false
			}
		}
		return true
	}
	for right < len(s) {
		MapWin[s[right]]++
		if compare(MapAim, MapWin, t) {
			break
		}
		right++
	}
	if right >= len(s) {
		return ""
	}
	var res [2]int
	res[0], res[1] = left, right
	for right < len(s) {
		for left <= right {
			if MapAim[s[left]] == 0 {
				left++
				continue
			}
			if MapWin[s[left]]-1 >= MapAim[s[left]] {
				MapWin[s[left]]--
				left++
			} else {
				if (right - left) < res[1]-res[0] {
					res[0], res[1] = left, right
				}
				MapWin[s[left]]--
				left++
				break
			}

		}

		for right++; right < len(s); right++ {
			MapWin[s[right]]++
			if s[right] == s[left-1] {
				break
			}
		}

	}

	return s[res[0] : res[1]+1]
}

// @lc code=end
