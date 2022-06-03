/*
 * @lc app=leetcode.cn id=606 lang=golang
 *
 * [606] 根据二叉树创建字符串
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package Solutions

import (
	"strconv"
	"strings"
)

func tree2str(root *TreeNode) string {
	var vlr func(*TreeNode, *strings.Builder)
	vlr = func(tn *TreeNode, b *strings.Builder) {

		b.WriteString(strconv.Itoa(tn.Val))
		if tn.Left != nil {
			b.WriteByte('(')
			vlr(tn.Left, b)
			b.WriteByte(')')
		}
		if tn.Right != nil {
			if tn.Left == nil {
				b.WriteString("()")
			}
			b.WriteByte('(')
			vlr(tn.Right, b)
			b.WriteByte(')')
		}

	}
	if root == nil {
		return ""
	}
	var sb = &strings.Builder{}
	vlr(root, sb)
	return sb.String()
}

// @lc code=end
