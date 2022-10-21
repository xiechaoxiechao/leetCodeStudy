/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */
package Solutions

// @lc code=start
type MyQueue struct {
	que []int
}

func Constructor__() MyQueue {
	m := MyQueue{}
	m.que = make([]int, 0, 8)
	return m

}

func (this *MyQueue) Push(x int) {
	this.que = append(this.que, x)
}

func (this *MyQueue) Pop() int {
	if !this.Empty() {
		t := this.que[0]
		this.que = this.que[1:]
		return t
	}
	panic("")

}

func (this *MyQueue) Peek() int {
	if !this.Empty() {
		return this.que[0]

	}
	panic("")
}

func (this *MyQueue) Empty() bool {
	if len(this.que) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
