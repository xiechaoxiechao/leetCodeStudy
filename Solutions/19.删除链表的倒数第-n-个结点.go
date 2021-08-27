package Solutions

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := make([]*ListNode, 0, 1)
	p = append(p, head)
	var length = 0
	var aim int
	for now := head; now != nil; now = now.Next {
		p = append(p, now.Next)
		length++
	}
	aim = length - n
	if aim == 0 {
		head = head.Next
		return head
	} else {
		p[aim-1].Next = p[aim].Next
		return head
	}

}

// @lc code=end
