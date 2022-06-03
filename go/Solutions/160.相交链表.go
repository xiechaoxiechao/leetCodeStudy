/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package Solutions

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var flagA bool
	var flagB bool
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			if flagA {
				p1 = headA
			} else {
				p1 = headB
			}
			flagA = !flagA
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			if flagB {
				p2 = headB
			} else {
				p2 = headA
			}
			flagB = !flagB
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

// @lc code=end
