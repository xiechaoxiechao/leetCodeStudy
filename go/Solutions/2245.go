package Solutions

func maxTrailingZeros(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	var fiv = make([][][2]int, m)
	var two = make([][][2]int, m)
	for i := 0; i < m; i++ {
		fiv[i] = make([][2]int, n)
		two[i] = make([][2]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fivNum := 0
			twoNum := 0
			t := grid[i][j]
			for t%2 == 0 {
				t /= 2
				twoNum++

			}
			t = grid[i][j]
			for t%5 == 0 {
				t /= 5
				fivNum++
			}
			if j > 0 {
				fiv[i][j][0] = fiv[i][j-1][0] + fivNum
				two[i][j][0] = two[i][j-1][0] + twoNum
			} else {
				fiv[i][j][0] = fivNum
				two[i][j][0] = twoNum
			}
			if i > 0 {
				fiv[i][j][1] = fiv[i-1][j][1] + fivNum
				two[i][j][1] = two[i-1][j][1] + twoNum
			} else {
				fiv[i][j][1] = fivNum
				two[i][j][1] = twoNum
			}
		}
	}
	var ans = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//left up
			t := 0
			f := 0
			if j > 0 {
				t = two[i][j-1][0] + two[i][j][1]
				f = fiv[i][j-1][0] + fiv[i][j][1]
			} else {
				t = two[i][j][1]
				f = fiv[i][j][1]
			}
			a := max_max(f, t)
			if a > ans {
				ans = a
			}
			//left down
			t = 0
			f = 0
			t = two[i][j][0] + two[m-1][j][1] - two[i][j][1]
			f = fiv[i][j][0] + fiv[m-1][j][1] - fiv[i][j][1]
			a = max_max(f, t)
			if a > ans {
				ans = a
			}
			//right up
			t = two[i][j][1] + two[i][n-1][0] - two[i][j][0]
			f = fiv[i][j][1] + fiv[i][n-1][0] - fiv[i][j][0]
			a = max_max(f, t)
			if a > ans {
				ans = a
			}
			//right down
			if j > 0 {
				t = two[i][n-1][0] - two[i][j-1][0] + two[m-1][j][1] - two[i][j][1]
				f = fiv[i][n-1][0] - fiv[i][j-1][0] + fiv[m-1][j][1] - fiv[i][j][1]
			} else {
				t = two[i][n-1][0] + two[m-1][j][1] - two[i][j][1]
				f = fiv[i][n-1][0] + fiv[m-1][j][1] - fiv[i][j][1]
			}
			a = max_max(f, t)
			if a > ans {
				ans = a
			}

		}
	}
	return ans
}

func max_max(i, j int) int {
	if i > j {
		return j
	}
	return i
}
