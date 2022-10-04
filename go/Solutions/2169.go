package Solutions

func countOperations(num1 int, num2 int) int {
	for i := 0; ; i++ {
		if num1 == 0 || num2 == 0 {
			return i
		}
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
	}
}