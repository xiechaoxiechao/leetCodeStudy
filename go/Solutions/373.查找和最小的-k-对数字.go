/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的 K 对数字
 */
package Solutions

import "container/heap"

type pq struct {
	data  *[][2]int
	nums1 []int
	nums2 []int
}

func (p pq) Len() int {
	return len(*(p.data))
}
func (p pq) Swap(i, j int) {
	(*(p.data))[i], (*(p.data))[j] = (*(p.data))[j], (*(p.data))[i]
}
func (p pq) Less(i, j int) bool {
	return p.nums1[(*(p.data))[i][0]]+p.nums2[(*(p.data))[i][1]] < p.nums1[(*(p.data))[j][0]]+p.nums2[(*(p.data))[j][1]]
}

func (p *pq) Push(v interface{}) {
	v1 := v.([2]int)
	(*(p.data)) = append((*(p.data)), v1)
}
func (p *pq) Pop() interface{} {
	item := (*(p.data))[len((*(p.data)))-1]
	(*(p.data)) = (*(p.data))[:len((*(p.data)))-1]
	return item
}

// @lc code=start
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var p *pq = new(pq)
	t := make([][2]int, 0, 8)
	p.data = &t
	p.nums1 = nums1
	p.nums2 = nums2
	heap.Init(p)
	for i := 0; i < len(nums1); i++ {
		heap.Push(p, [2]int{i, 0})
	}

	ans := make([][]int, 0, 8)
	for i := 0; i < k; i++ {
		if len(*p.data) == 0 {
			break
		}
		item := heap.Pop(p).([2]int)
		ans = append(ans, []int{nums1[item[0]], nums2[item[1]]})
		if item[1]+1 < len(nums2) {
			heap.Push(p, [2]int{item[0], item[1] + 1})
		}

	}
	return ans

}

// @lc code=end
