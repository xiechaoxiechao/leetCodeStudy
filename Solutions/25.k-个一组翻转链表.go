package Solutions

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	stack := make([]*ListNode, k)
	i := 0
	var now *ListNode
	c := 0
	var nowH *ListNode
	nowE := head
	for now = head; now != nil; now = now.Next {
		stack[i] = now
		i++
		c++
		if c == k {
			c = 0
			nowE = now.Next
			for i := k - 1; i > 0; i-- {
				stack[i].Next = stack[i-1]
			}
			if nowH == nil {
				nowH = stack[k-1]
				head = nowH
			} else {

				nowH.Next = stack[k-1]
			}
			stack[0].Next = nowE
			now = stack[0]
			nowH = now
			i = 0

		}
	}

	return head
}

// @lc code=end
