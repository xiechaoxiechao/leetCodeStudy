/*
 * @lc app=leetcode.cn id=668 lang=c
 *
 * [668] 乘法表中第k小的数
 */

// @lc code=start

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int findKthNumber(int m, int n, int k) {
  int left = 1, right = m * n;
  while (left < right) {
    int x = left + (right - left) / 2;
    int count = x / n * n;
    for (int i = x / n + 1; i <= m; ++i) {
      count += x / i;
    }
    if (count >= k) {
      right = x;
    } else {
      left = x + 1;
    }
  }
  return left;
}

int main() { findKthNumber(3, 3, 5); }
// @lc code=end
