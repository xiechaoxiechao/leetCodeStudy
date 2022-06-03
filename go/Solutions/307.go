package Solutions

type NumArray struct {
	_dat  []int
	_tree []int
}

func Constructor(nums []int) NumArray {
	var n = NumArray{}
	n._dat = nums[:]
	h := 0
	for len(nums) > (1 << h) {
		h++
	}

	n._tree = make([]int, 1<<(h+1)-1)
	off := 1<<h - 1
	copy(n._tree[off:], nums[:])
	return n

}

func (this *NumArray) Update(index int, val int) {
	this._dat[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	var sum = 0
	for i := left; i <= right; i++ {
		sum += this._dat[i]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
