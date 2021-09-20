package Solutions

/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	now, last := head.Next, head
	for ; now != nil; now = now.Next {
		if now.Val != last.Val {
			last.Next = now
			last = now
		}
	}
	last.Next = now

	return head
}

// @lc code=end
