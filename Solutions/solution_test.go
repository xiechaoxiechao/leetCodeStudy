package Solutions

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {

}

func BenchmarkCombine(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		fmt.Println(combine(4, 2))
		// TODO: Your Code Here
	}
}

func BenchmarkExist(b *testing.B) {
	a := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(a, "ABCCED"))
}

func TestSearch(t *testing.T) {
	a := []int{1, 3}
	t.Log(search1(a, 0))

}
