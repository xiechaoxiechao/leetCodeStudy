/*
 * @lc app=leetcode.cn id=2 lang=javascript
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
 var addTwoNumbers = function (l1, l2) {
    var k = l1;
    mul = 1n;
    var add=0n;
    var mid;
    do {
      mid = BigInt(k.val);
      add += mid * mul;
      mul *= 10n;
      k = k.next;
    } while (k!= null)
    mul=1n;
    k=l2;
    do {
      mid = BigInt(k.val);
      add += mid * mul;
      mul *= 10n;
      k = k.next;
    } while (k!= null)
    var str=add.toString().split('');
    var res=new ListNode(str[0]);
    var m;
    for(var i=1,n=str.length;i<n;i++){
      m=new ListNode(str[i]);
      m.next=res;
      res=m;
    }
    return res;
  
  
  };
// @lc code=end

