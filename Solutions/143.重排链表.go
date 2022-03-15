/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start

package Solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	var len = 0
	for now := head; now != nil; now = now.Next {
		len++
	}
	mid := (len + 1) / 2
	head2 := head
	for i := 1; i < mid; i++ {
		head2 = head2.Next
	}
	var last *ListNode = nil
	var tem *ListNode = nil
	for now := head2.Next; now != nil; now = tem {
		tem = now.Next
		now.Next = last
		last = now
	}
	head2.Next = nil
	t := head
	now1, now2 := head.Next, last
	for ; now1 != nil; now1, now2 = now1.Next, tem {
		t.Next = now2
		tem = now2.Next
		now2.Next = now1
		t = now1
	}
	if now2 != nil {
		t.Next = now2
	}

}

// @lc code=end
