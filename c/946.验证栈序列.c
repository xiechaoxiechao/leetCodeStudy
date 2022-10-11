/*
 * @lc app=leetcode.cn id=946 lang=c
 *
 * [946] 验证栈序列
 */

// @lc code=start
#include <stdbool.h>

bool validateStackSequences(int *pushed, int pushedSize, int *popped,
                            int poppedSize) {
  int *stack = malloc(sizeof(int) * pushedSize);
  int stackTop = -1;
  int j = 0;
  for (int i = 0; i < pushedSize; i++) {
    stackTop++;
    stack[stackTop] = pushed[i];
    while (stackTop >= 0 && stack[stackTop] == popped[j]) {
      stackTop--;
      j++;
    }
  }
  while (stackTop >= 0) {
    if (stack[stackTop] != popped[j]) {
      return false;
    }
    stackTop--;
    j++;
  }
  return true;
}

bool validateStackSequences(int *pushed, int pushedSize, int *popped,
                            int poppedSize) {
  bool *status = calloc(pushedSize, sizeof(bool));
  memset(status, 0, sizeof(bool) * pushedSize);
  int begin = 0;
  for (int i = 0; i < pushedSize; i++) {
    if (status[i]) {
      continue;
    }
    status[i] = true;
    int j = begin;
    while (j < pushedSize) {
      if (pushed[j] == popped[i]) {
        break;
      }
      j++;
    }
    for (int l1 = j - 1, l2 = i + 1; l1 >= begin; l1--) {
      while (l2 < pushedSize && popped[l2] != pushed[l1]) {
        if (status[l2]) {
          return false;
        }
        l2++;
      }
      if (l2 == pushedSize) {
        return false;
      }
      status[l2] = true;
      l2++;
    }
    begin = j + 1;
  }
  return true;
}

// @lc code=end
