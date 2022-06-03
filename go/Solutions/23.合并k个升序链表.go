package Solutions

import "sort"

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	b := make([]int, 0, 1)
	hs := make(map[int][]*ListNode)
	var now *ListNode
	for i := 0; i < length; i++ {
		for now = lists[i]; now != nil; now = now.Next {
			value, ok := hs[now.Val]
			if ok == false {
				hs[now.Val] = []*ListNode{now}
			} else {
				hs[now.Val] = append(value, now)
			}
			b = append(b, now.Val)
		}
	}
	sort.Ints(b)
	length2 := len(b)
	var head *ListNode
	if length2 == 0 {
		return nil
	}
	for _, i := range hs[b[0]] {
		if head == nil {
			head = i
			now = i
		} else {
			now.Next = i
			now = now.Next
		}
	}
	if length2 == 1 {
		return head
	}
	i := 0
	for ; i < length2-1 && b[i] == b[i+1]; i++ {
	}
	i++
	for ; i < length2; i++ {
		for _, i := range hs[b[i]] {
			now.Next = i
			now = now.Next
		}
		for ; i < length2-1 && b[i] == b[i+1]; i++ {
		}
	}
	return head
}

// @lc code=end
