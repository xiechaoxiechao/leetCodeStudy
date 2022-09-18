package Solutions

import "math"

func minSteps(n int) int {
	num := 1
	mem := 1
	if n == 1 {
		return 0
	}
	return dfs(1, num, n, mem)

}

func dfs(step int, num int, aim int, mem int) int {

	step1, step2 := math.MaxInt32, math.MaxInt32
	if num+mem == aim {
		return step + 1
	} else if num+mem < aim {
		step1 = dfs(step+1, num+mem, aim, mem)
	} else {
		return math.MaxInt32
	}

	if mem != num {
		step2 = dfs(step+1, num, aim, num)
	}
	if step1 < step2 {
		return step1
	}
	return step2

}
