package Solutions

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
var numMap map[byte][]string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	numMap = make(map[byte][]string)
	numMap[49] = []string{}
	numMap[50] = []string{"a", "b", "c"}
	numMap[51] = []string{"d", "e", "f"}
	numMap[52] = []string{"g", "h", "i"}
	numMap[53] = []string{"j", "k", "l"}
	numMap[54] = []string{"m", "n", "o"}
	numMap[55] = []string{"p", "q", "r", "s"}
	numMap[56] = []string{"t", "u", "v"}
	numMap[57] = []string{"w", "x", "y", "z"}
	if len(digits) == 1 {
		return numMap[digits[0]]
	}
	j := 1
	return turn(j, digits, numMap[digits[0]])
}
func turn(a int, c string, n []string) []string {
	h := make([]string, 0)
	j := a + 1

	for w, e := len(n), 0; e < w; e++ {
		for k, i := len(numMap[c[j-1]]), 0; i < k; i++ {
			h = append(h, n[e]+numMap[c[j-1]][i])
		}
	}
	if j == len(c) {
		return h
	}

	return turn(j, c, h)

}

// @lc code=end
