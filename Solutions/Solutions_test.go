package Solutions

import (
	"fmt"
	"testing"
)

func TestMinDistance(t *testing.T) {
	if minDistance("horse", "ros") != 3 {
		t.Error("error")
	}
}

// func BenchmarkMindistance(b *testing.B) {
// 	// for i := 0; i < 2; i++ {
// 	minDistance("horse", "ros")
// 	// }
// }
func BenchmarkSearchMartix(b *testing.B) {
	var a = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix(a, 13))
}

func BenchmarkMainWindow(b *testing.B) {
	fmt.Println(minWindow("a", "aa"))
}
