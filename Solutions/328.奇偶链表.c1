/*
 * @lc app=leetcode.cn id=328 lang=c
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
#include<stdio.h>
#include<stdbool.h>
struct ListNode {
    int val;
    struct ListNode *next;
};
struct ListNode* oddEvenList(struct ListNode* head){
    struct ListNode * oddHeade;
    struct ListNode * oddTail;
    struct ListNode * now =head;
    if(now ==NULL||now->next==NULL){
        return head;
    }
    oddHeade=now->next;
    oddTail=oddHeade;
    if(oddHeade->next==NULL){
        return head;
    }
    now->next=oddHeade->next;
    oddTail->next=NULL;
    now=now->next;
    while (true){
        if (now->next==NULL){
            break;
        }
        oddTail->next=now->next;
        oddTail=oddTail->next;
        now->next=oddTail->next;
        oddTail->next=NULL;
        if(now->next==NULL){
            break;
        }
        now=now->next;
    }
    now->next=oddHeade;
    return head;
}
int main(){
    struct ListNode l1,l2,l3,l4,l5;
    l1.val=1;
    l1.next=&l2;
    l2.val=2;
    l2.next=&l3;
    l3.val=3;
    l3.next=&l4;
    l4.val=4;
    l4.next=&l5;
    l5.val=5;
    l5.next=NULL;
    oddEvenList(&l1);
    printf("%d",l1.next->next->next->next->val);
}
// @lc code=end

