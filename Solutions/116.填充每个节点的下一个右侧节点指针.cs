using System.Collections.Generic;
using System.Collections.ObjectModel;
/*
 * @lc app=leetcode.cn id=116 lang=csharp
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start

public class Node {
    public int val;
    public Node left;
    public Node right;
    public Node next;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _left, Node _right, Node _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
}
public partial class Solution {
    public Node Connect(Node root) {
        if (root ==null){
            return root;
        }
        var items=new List<Node>();
        var i=0;
        items.Add(root);
        var depth=1;
        var off=0;
        while(i<items.Count){
            if (items[i].left != null){
                items.Add(items[i].left);
            }
            if (items[i].right != null){
                items.Add(items[i].right);
            }
            if (i==off){
                off+=1<<depth;
                depth++;
                i++;
                continue;
            }
            items[i].next=items[i+1];
            i++;

            
        } 
        return root;
    }
}

// @lc code=end

