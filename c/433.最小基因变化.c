/*
 * @lc app=leetcode.cn id=433 lang=c
 *
 * [433] 最小基因变化
 */

// @lc code=start
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
bool isEffect(const char *a, const char *b);
int minMutation(char *start, char *end, char **bank, int bankSize) {
  int endI = -1;
  for (int i = 0; i < bankSize; i++) {
    if (memcmp(bank[i], end, 8) == 0) {
      endI = i;
      break;
    }
  }
  if (endI == -1) {
    return -1;
  }
  if (memcmp(start, end, 8) == 0) {
    return 0;
  }
  int beginI = 0;
  for (int i = 0; i < bankSize; i++) {
    if (isEffect(start, bank[i])) {
      beginI |= 1 << i;
    }
  }
  int reachable[bankSize][bankSize];
  memset(reachable, 0, sizeof(int) * bankSize * bankSize);
  for (int i = 0; i < bankSize; i++) {
    for (int j = i + 1; j < bankSize; j++) {
      if (isEffect(bank[i], bank[j])) {
        reachable[i][j] = 1;
        reachable[j][i] = 1;
      }
    }
  }
  int reach = 0;
  reach = 1 << endI;
  if (isEffect(start, end)) {
    return 1;
  }
  for (int distance = 1; true; distance++) {
    int reach2 = 0;
    for (int i = 0; i < bankSize; i++) {
      if ((reach & 1 << i) != 0) {
        for (int j = 0; j < bankSize; j++) {
          if (reachable[i][j] == 1) {
            reach2 |= (1 << j);
          }
        }
      }
    }
    if ((beginI & reach2) != 0) {
      return distance + 1;
    }
    if (reach2 == 0) {
      return -1;
    }
    reach = reach2;
  }

  return 0;
}
bool isEffect(const char *a, const char *b) {
  int dif = 0;
  for (int i = 0; i < 8; i++) {
    if (a[i] != b[i]) {
      dif++;
    }
  }
  if (dif <= 1) {
    return true;
  }
  return false;
}
// @lc code=end
