/*
 * @lc app=leetcode.cn id=1865 lang=golang
 *
 * [1865] 找出和为指定值的下标对
 */
package Solutions

// @lc code=start
type FindSumPairs struct {
	nodes1 []int
	nodes2 map[int]int
	ind    []int
}

func Constructor___(nums1 []int, nums2 []int) FindSumPairs {
	c := FindSumPairs{}
	c.nodes1 = nums1
	c.nodes2 = make(map[int]int, len(nums2))
	c.ind = nums2
	for i := 0; i < len(nums2); i++ {
		c.nodes2[nums2[i]]++
	}
	return c
}

func (this *FindSumPairs) Add(index int, val int) {
	t := this.ind[index]
	this.ind[index] += val
	this.nodes2[t] -= 1
	this.nodes2[this.ind[index]] += 1
}

func (this *FindSumPairs) Count(tot int) int {
	len1 := len(this.nodes1)
	ans := 0
	for i := 0; i < len1; i++ {
		if val, ok := this.nodes2[tot-this.nodes1[i]]; ok {
			ans += val
		}
	}
	return ans
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
// @lc code=end
