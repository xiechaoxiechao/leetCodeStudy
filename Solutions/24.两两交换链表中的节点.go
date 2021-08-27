package Solutions

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func swapPairs(head *ListNode) *ListNode {
	Nhead := &ListNode{0, head}
	var last, ch, ch2 *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	for last = Nhead; true; {
		if last.Next.Next == nil {
			break
		}
		ch = last.Next.Next.Next
		ch2 = last.Next.Next
		last.Next.Next.Next = last.Next
		last.Next = ch2
		last.Next.Next.Next = ch
		if ch != nil {
			last = last.Next.Next
		} else {
			break
		}

	}
	return Nhead.Next

}

// @lc code=end
