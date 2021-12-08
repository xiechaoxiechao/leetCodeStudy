using System.Collections.Generic;
using System.Reflection.Metadata;
/*
 * @lc app=leetcode.cn id=82 lang=csharp
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start

public partial class Solutions
{
    public ListNode DeleteDuplicates(ListNode head)
    {
        ListNode now;
        ListNode last;
        ListNode tem;
        var res = new ListNode(0, head);
        for (last = res; true;)
        {
            if (last.next == null)
            {
                return res.next;
            }
            tem = last.next;
            now = tem.next;
            if (now == null)
            {
                return res.next;
            }
            for (; now != null && now.val == tem.val; now = now.next) { }
            if (now == null)
            {
                last.next = null;
                return res.next;
            }
            if (now == tem.next)
            {
                last = tem;
            }
            else
            {
                last.next = now;
            }

        }
    }
}
// @lc code=end

