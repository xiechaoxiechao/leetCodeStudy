/*
 * @lc app=leetcode.cn id=450 lang=c
 *
 * [450] 删除二叉搜索树中的节点
 */

// @lc code=start
/**
 * Definition for a binary tree node.

 */
#include<stdio.h>
 struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
 };

void removeNode(int key,struct TreeNode * root,struct TreeNode *parent);
struct TreeNode* deleteNode(struct TreeNode* root, int key){
    if(root==NULL){
        return NULL;
    }
    struct TreeNode* aim=root;
    struct TreeNode * last=NULL;
    int pos=0;
    while(aim->val!=key){
        last=aim;
        if(aim->val<key){
            if(aim->right==NULL){
                aim=NULL;
                break;
            }
            aim=aim->right;
            pos=1;
        }else{
            if(aim->left==NULL){
                aim=NULL;
                break;
            }
            aim=aim->left;
            pos=0;
        }
        
    }
    if(aim==NULL){
        return root;
    }
    if (aim==root&&aim->left==NULL&&aim->right==NULL){
        return NULL;
    }
    removeNode(pos,aim,last);
    return root;
     
   
}
void removeNode(int key,struct TreeNode * root,struct TreeNode *parent){
    struct TreeNode * tem=NULL;
    struct TreeNode * par=root;
    int o=-1;
    if (root->left!=NULL){
        tem=root->left;
        o=0;
        if(tem->right!=NULL){
             par=tem;
            tem=tem->right;
            o=1;
        }
        while(tem->right!=NULL){
            par=tem;
            tem=tem->right;
        }
        root->val=tem->val;
        removeNode(o,tem,par);
        return;
    }
     if(root->right!=NULL){
        tem=root->right;
         o=1;
        if(tem->left!=NULL){
             par=tem;
            tem=tem->left;
            o=0;
        }
        while(tem->left!=NULL){
               par=tem;
            tem=tem->left;
        }
        
        root->val=tem->val;
        removeNode(o,tem,par);
        return; 
    }
    
   
    if(key==1){
        parent->right=NULL;
    }else{
        parent->left=NULL;
    }
}