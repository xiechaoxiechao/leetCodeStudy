package Solutions

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var now1, now2, head, now *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		head = l1
		now1 = l1.Next
		now2 = l2

	} else {
		head = l2
		now2 = l2.Next
		now1 = l1
	}
	now = head
	for now1 != nil && now2 != nil {
		if now1.Val < now2.Val {
			now.Next = now1
			now = now.Next
			now1 = now1.Next
		} else {
			now.Next = now2
			now = now.Next
			now2 = now2.Next
		}

	}
	if now1 == nil {
		now.Next = now2
	} else {
		now.Next = now1
	}
	return head

}

// @lc code=end
