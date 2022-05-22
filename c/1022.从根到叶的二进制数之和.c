/*
 * @lc app=leetcode.cn id=1022 lang=c
 *
 * [1022] 从根到叶的二进制数之和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
#include<stdio.h>
#include<stdbool.h>
struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};
void dfs(const struct TreeNode*,int*,int*);
int sumRootToLeaf(struct TreeNode* root){
     if(root==NULL){
        return 0;
    }
    int sum=0;
    int value =0;
    dfs(root,&sum,&value);
    return sum;
}
void dfs(const struct TreeNode * root,int *sum,int * value){
    *value<<=1;
    *value|=root->val;
    bool flag=false;
    if(root->left!=NULL){
        dfs(root->left,sum,value);
    }else{
        flag=true;
    }
    if(root->right!=NULL){
        dfs(root->right,sum,value);
    }else if (flag){
        *sum+=*value;
    }
    *value>>=1;
}
// @lc code=end

