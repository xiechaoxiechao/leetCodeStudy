/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1 = make([]int, 0, 100)
	var stack2 = make([]int, 0, 100)
	for now := l1; now != nil; now = now.Next {
		stack1 = append(stack1, now.Val)
	}
	for now := l2; now != nil; now = now.Next {
		stack2 = append(stack2, now.Val)
	}
	var len1 = len(stack1) - 1
	var len2 = len(stack2) - 1
	var carry = 0
	var last *ListNode = nil
	for len1 >= 0 && len2 >= 0 {
		tem := stack1[len1] + stack2[len2] + carry
		carry = 0
		len1--
		len2--
		if tem >= 10 {
			tem -= 10
			carry = 1
		}
		var temNode = ListNode{tem, last}
		last = &temNode
	}
	for len1 >= 0 {
		tem := stack1[len1] + carry
		len1--
		carry = 0
		if tem >= 10 {
			tem -= 10
			carry = 1
		}
		var temNode = ListNode{tem, last}
		last = &temNode
	}
	for len2 >= 0 {
		tem := stack2[len2] + carry
		len2--
		carry = 0
		if tem >= 10 {
			tem -= 10
			carry = 1
		}
		var temNode = ListNode{tem, last}
		last = &temNode
	}
	if carry == 1 {
		var temNode = ListNode{1, last}
		last = &temNode
	}
	return last
}

// @lc code=end
