/*
 * @lc app=leetcode.cn id=1514 lang=golang
 *
 * [1514] 概率最大的路径
 */
package Solutions

import (
	"container/heap"
)

// @lc code=start

type __node struct {
	index int
	prob  float64
}
type NodeHeap []__node

func (h NodeHeap) Len() int {
	return len(h)
}

func (n NodeHeap) Less(i, j int) bool {
	return n[i].prob < n[j].prob
}

func (n *NodeHeap) Push(x interface{}) {
	*n = append(*n, x.(__node))
}
func (n *NodeHeap) Pop() interface{} {
	old := (*n)
	lenOld := len(old)
	x := old[lenOld-1]
	(*n) = old[0 : lenOld-1]
	return x
}

func (n NodeHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	var mem = make([]float64, n)
	var brod = make([][]int, n)
	var prob = make([][]float64, n)
	var nodeHeap NodeHeap = make([]__node, 0, n)

	for i := 0; i < n; i++ {
		brod[i] = make([]int, 0, 8)
		prob[i] = make([]float64, 0, 8)
	}
	heap.Init(&nodeHeap)
	for i := 0; i < len(edges); i++ {
		brod[edges[i][0]] = append(brod[edges[i][0]], edges[i][1])
		brod[edges[i][1]] = append(brod[edges[i][1]], edges[i][0])
		prob[edges[i][0]] = append(prob[edges[i][0]], succProb[i])
		prob[edges[i][1]] = append(prob[edges[i][1]], succProb[i])
	}
	heap.Push(&nodeHeap, __node{start, 1})

	for len(nodeHeap) != 0 {
		current := heap.Pop(&nodeHeap).(__node)
		if current.prob <= mem[current.index] {
			continue
		}
		mem[current.index] = current.prob
		for i, v := range brod[current.index] {
			pp := prob[current.index][i] * current.prob
			if pp < mem[v] {
				heap.Push(&nodeHeap, __node{v, prob[current.index][i] * current.prob})
			}

		}

	}
	return mem[end]

}

// @lc code=end
