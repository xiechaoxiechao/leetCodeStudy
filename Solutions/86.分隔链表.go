package Solutions

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var leftHead ListNode
	var rightHead ListNode
	leftHead.Next = head
	rightTail := &rightHead
	pre := &leftHead
	for now := head; now != nil; now = now.Next {
		if now.Val >= x {
			pre.Next = now.Next
			rightTail.Next = now
			rightTail = now
		} else {
			pre = now
		}
	}
	pre.Next = rightHead.Next
	rightTail.Next = nil
	return leftHead.Next

}

// @lc code=end
