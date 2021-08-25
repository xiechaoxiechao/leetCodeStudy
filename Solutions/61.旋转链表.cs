using System.Collections.Generic;
using System;
/*
 * @lc app=leetcode.cn id=61 lang=csharp
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class ListNode
{
    public int val;
    public ListNode next;
    public ListNode(int val = 0, ListNode next = null)
    {
        this.val = val;
        this.next = next;
    }
}
public partial class Solution
{
    public ListNode RotateRight(ListNode head, int k)
    {
        if(head==null){
            return head;
        }
        var tem = new Queue<ListNode>();
        ListNode i;
        for(i=head;i!=null;i=i.next){
            tem.Enqueue(i);
        }
        var nodeArray=tem.ToArray();
        Array.Reverse(nodeArray);
        tem=new Queue<ListNode>(nodeArray);
        k=k-(k/tem.Count)*tem.Count;
        for(int j=0;j<k;++j){
            i=tem.Dequeue();
            tem.Enqueue(i);
        }
        nodeArray=tem.ToArray();
        Array.Reverse(nodeArray);
        for(int j=0;j<k&&j<tem.Count;++j){
            if(j==tem.Count-1){
                nodeArray[j].next=null;
                break;
            }else{
                nodeArray[j].next=nodeArray[j+1];
            }

        }
        nodeArray[nodeArray.Length-1].next=null;
        return nodeArray[0];
    }
}
// @lc code=end

