/*
 * @lc app=leetcode.cn id=2185 lang=c
 *
 * [2185] 统计包含给定前缀的字符串
 */

// @lc code=start
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int prefixCount(char **words, int wordsSize, char *pref) {
  int ans = 0;
  int len = strlen(pref);
  for (int i = 0; i < wordsSize; i++) {
    if (strlen(words[i]) < len) {
      continue;
    }
    if (memcmp(words[i], pref, len) == 0) {
      ++ans;
    }
  }
  return ans;
}
// @lc code=end
