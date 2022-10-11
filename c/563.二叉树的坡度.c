/*
 * @lc app=leetcode.cn id=563 lang=c
 *
 * [563] 二叉树的坡度
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
#include <stdio.h>
#include <stdlib.h>
struct TreeNode {
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};
int _dfs(struct TreeNode *root, int *result);

int findTilt(struct TreeNode *root) {
  int result;
  _dfs(root, &result);
  return result;
}

int _dfs(struct TreeNode *root, int *result) {
  int left = root->left == NULL ? 0 : _dfs(root->left, result);
  int right = root->right == NULL ? 0 : _dfs(root->right, result);
  result += abs(left - right);
  return left + right + root->val;
}
// @lc code=end
