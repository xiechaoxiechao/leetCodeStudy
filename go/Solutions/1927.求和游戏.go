/*
 * @lc app=leetcode.cn id=1927 lang=golang
 *
 * [1927] 求和游戏
 */
package Solutions

// @lc code=start
func sumGame(num string) bool {
	leftSum, leftCount := sum(num[:len(num)/2])
	rightSum, rightCount := sum(num[len(num)/2:])
	isAlice := true
	for rightCount > 0 || leftCount > 0 {
		if isAlice {
			isAlice = false
			if leftSum > rightSum {
				if leftCount > 0 {
					leftSum += 9
					leftCount--
				} else {
					rightCount--
				}
			} else if leftSum == rightSum {
				if leftCount > 0 {
					leftSum += 9
					leftCount--
				} else {
					rightSum += 9
					rightCount--
				}
			} else {
				if rightCount > 0 {
					rightSum += 9
					rightCount--
				} else {
					leftCount--
				}
			}
		} else {
			isAlice = true
			if leftSum < rightSum {
				if leftCount > 0 {
					if rightSum-leftSum > 9 {
						leftSum += 9
						leftCount--
					} else {
						leftSum = rightSum
						leftCount--
					}
				} else {
					rightCount--
				}
			} else if leftSum == rightSum {
				if leftCount > 0 {
					leftCount--
				} else {
					rightCount--
				}
			} else {
				if rightCount > 0 {
					if leftSum-rightSum > 9 {
						rightSum += 9
						rightCount--
					} else {
						rightSum = leftSum
						rightCount--
					}
				}
			}
		}
	}

	return leftSum != rightSum

}
func sum(s string) (int, int) {
	sum := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '?' {
			sum += int(s[i] - '0')
		} else {
			count++
		}

	}
	return sum, count
}

// @lc code=end
