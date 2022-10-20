/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
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
func removeElements(head *ListNode, val int) *ListNode {
	parent := &ListNode{
		0,
		head,
	}
	head_h := parent
	for parent.Next != nil {
		if parent.Next.Val == val {
			parent.Next = parent.Next.Next
		} else {
			parent = parent.Next
		}

	}
	return head_h.Next
}

// @lc code=end
