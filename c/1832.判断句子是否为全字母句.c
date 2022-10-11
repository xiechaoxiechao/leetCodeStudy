/*
 * @lc app=leetcode.cn id=1832 lang=c
 *
 * [1832] 判断句子是否为全字母句
 */

// @lc code=start
#include <stdbool.h>
#include <string.h>

bool checkIfPangram(char *sentence) {
  int b = 0;
  int n = strlen(sentence);
  for (int i = 0; i < n; i++) {
    b |= (1 << (sentence[i] - 'a'));
  }
  return b == 0b0000011111111111111111111111111;
}
// @lc code=end
