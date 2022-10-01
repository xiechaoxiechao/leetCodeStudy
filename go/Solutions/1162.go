package Solutions

func maxDistance(grid [][]int) int {
	const MAX int = 1<<32 - 1
	var n = len(grid)
	var distance = make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, n)
	}
	var max = -1
	var que = make([][2]int, 0, 32)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				distance[i][j] = 0
				que = append(que, [2]int{i - 1, j}, [2]int{i + 1, j}, [2]int{i, j - 1}, [2]int{i, j + 1})
			} else {
				distance[i][j] = MAX
			}
		}
	}
	for ind := 0; ind < len(que); ind++ {

		i, j := que[ind][0], que[ind][1]
		if i < 0 || j < 0 || i >= n || j >= n {
			continue
		}
		min := MAX
		if i-1 >= 0 {

			if min > distance[i-1][j] {
				min = distance[i-1][j]
			}
		}
		if j-1 >= 0 {
			if min > distance[i][j-1] {
				min = distance[i][j-1]
			}
		}
		if i+1 < len(grid) {
			if min > distance[i+1][j] {
				min = distance[i+1][j]
			}
		}
		if j+1 < len(grid) {
			if min > distance[i][j+1] {
				min = distance[i][j+1]
			}
		}
		min++
		if min < distance[i][j] {
			distance[i][j] = min
			que = append(que, [2]int{i - 1, j}, [2]int{i + 1, j}, [2]int{i, j - 1}, [2]int{i, j + 1})
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if max < distance[i][j] && distance[i][j] != 0 {
				max = distance[i][j]
			}
		}
	}
    if max == MAX {
		return -1
	}
	return max

}
