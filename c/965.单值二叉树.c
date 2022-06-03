/*
 * @lc app=leetcode.cn id=965 lang=c
 *
 * [965] 单值二叉树
 */

// @lc code=start

 
struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

#include<stdbool.h>
#include<stdio.h>

bool isUnivalTree(struct TreeNode* root){
    if (root->left!=NULL){
        if(root->left->val!=root->val||!isUnivalTree(root->left)){
            return false;
        }
    }
    if (root->right!=NULL){
        if(root->right->val!=root->val||!isUnivalTree(root->right)){
            return false;
        }
    }
    return true;
}

// @lc code=end

