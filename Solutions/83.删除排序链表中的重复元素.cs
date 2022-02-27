using System.Collections.Generic;
/*
 * @lc app=leetcode.cn id=83 lang=csharp
 *
 * [83] 删除排序链表中的重复元素
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
public partial class Solution {
    public ListNode DeleteDuplicates(ListNode head) {
        var h = new ListNode(-101,head);
        for(ListNode now=h,next=h.next;next != null;next=next.next){
            if(now.val==next.val){
                now.next=next.next;
            }else{
                now=next;
            }  
        }
        return h.next;
    }
}
// @lc code=end

