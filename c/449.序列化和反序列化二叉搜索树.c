/*
 * @lc app=leetcode.cn id=449 lang=c
 *
 * [449] 序列化和反序列化二叉搜索树
 */

// @lc code=start
#include <stdio.h>
#include <stdlib.h>
struct TreeNode {
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};
/** Encodes a tree to a single string. */
int *se(int *c, struct TreeNode *node);
struct TreeNode *de_se(int *data, int begin, int end);
char *serialize(struct TreeNode *root) {
  int *buf = malloc(10000);
  int *end = se(buf, root);
  *end = -1;
  return (char *)buf;
}

/** Decodes your encoded data to tree. */
struct TreeNode *deserialize(char *data) {
  int *dat = (int *)data;
  int end = 0;
  for (; dat[end] != -1; end++) {
  }
  if (end == 0) {
    return NULL;
  }
  return de_se(dat, 0, end);
}
int *se(int *c, struct TreeNode *node) {
  if (node == NULL) {
    return c;
  }
  *c = node->val;
  c++;
  c = se(c, node->left);
  c = se(c, node->right);
  return c;
}
struct TreeNode *de_se(int *data, int begin, int end) {
  struct TreeNode *node = malloc(sizeof(struct TreeNode));
  node->val = data[begin];
  begin++;
  int i = 0;
  for (i = begin; i < end; i++) {
    if (data[i] > data[begin]) {
      break;
    }
  }

  if (i > begin) {
    node->left = de_se(data, begin, i);
  } else {
    node->left = NULL;
  }
  if (i < end) {
    node->right = de_se(data, i, end);
  } else {
    node->right = NULL;
  }
  return node;
}
// Your functions will be called as such:
// char* data = serialize(root);
// deserialize(data);
// @lc code=end
