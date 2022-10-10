/*
 * @lc app=leetcode.cn id=2074 lang=golang
 *
 * [2074] 反转偶数长度组的节点
 */
package Solutions

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseEvenLengthGroups(head *ListNode) *ListNode {
	last := &ListNode{}
	a := last
	last.Next = head
	for i, nowHead := 1, head; nowHead != nil; i++ {
		h, t, next := reserve(nowHead, i, 1)
		if h != nil {
			last.Next = h
			t.Next = next
		}
		last = t
		nowHead = next
	}
	return a.Next
}

func reserve(head *ListNode, aimLen int, depth int) (*ListNode, *ListNode, *ListNode) {
	if head.Next == nil || aimLen == depth {
		if depth%2 == 0 {
			return head, head, head.Next
		} else {
			return nil, head, head.Next
		}
	}
	h, t, next := reserve(head.Next, aimLen, depth+1)
	if h == nil {
		return h, t, next
	}
	t.Next = head
	return h, head, next

}

// @lc code=end
