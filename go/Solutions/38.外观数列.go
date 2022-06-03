package Solutions

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	res := make([]byte, 1)
	res[0] = '1'
	if n == 1 {
		return string(res)
	}
	var a byte
	var c byte
	count := 0
	tem := make([]byte, 0)
	for i := 2; i <= n; i++ {
		a = res[0]
		count = 0
		for _, c = range res {
			if c == a {
				count++
			} else {
				tem = append(tem, byte(count+48), a)
				a = c
				count = 1
			}
		}
		tem = append(tem, byte(count+48), a)
		res = tem
		tem = make([]byte, 0)
	}
	return string(res)
}

// @lc code=end
