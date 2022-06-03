package Solutions

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	a := make([][]string, 0, 1)
	a = append(a, []string{""}, []string{"()"})
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []string{"()"}
	}

	for i := 2; i <= n; i++ {
		b := make([]string, 0, 1)
		for p, q := 0, i-1; p <= i-1 && q >= 0; p, q = p+1, q-1 {
			for _, c := range a[p] {
				for _, d := range a[q] {
					b = append(b, "("+c+")"+d)
				}
			}
		}
		a = append(a, b)

	}
	//fmt.Println(a[3])
	return a[n]

}

// @lc code=end
