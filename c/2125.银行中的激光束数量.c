/*
 * @lc app=leetcode.cn id=2125 lang=c
 *
 * [2125] 银行中的激光束数量
 */

// @lc code=start

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int numberOfBeams(char **bank, int bankSize) {
  int ans = 0;
  int m = strlen(bank[0]) / sizeof(char);
  int *count = calloc(bankSize, sizeof(int));
  // memset(0,count,strlen(count));
  for (int i = 0; i < bankSize; i++) {
    for (int j = 0; j < m; j++) {
      if (bank[i][j] == '1') {
        count[i]++;
      }
    }
  }
  int last = -1;
  for (int i = 0; i < bankSize; i++) {
    if (count > 0) {
      if (last != -1) {
        ans += count[last] * count[i];
      }
      last = i;
    }
  }
  return ans;
}
// @lc code=end
