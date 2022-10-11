/*
 * @lc app=leetcode.cn id=467 lang=c
 *
 * [467] 环绕字符串中唯一的子字符串
 */

// @lc code=start

#include <stdio.h>
#include <stdlib.h>
int findSubstringInWraproundString(char *p) {
  int l = 0;
  int ans = 0;
  int i = 0;
  while (1) {
    char tem = p[i] + 1;
    if (tem == '{') {
      tem = 'a';
    }
    if (p[i + 1] != tem) {
      int t = i - l + 1;
      ans += (t + 1) * t / 2;
      l = i + 1;
    }
    i++;
    if (p[i] == '\0') {
      int t1 = i - l;
      ans += (t1 + 1) * t1 / 2;
      break;
    }
  }
  return ans;
}
int main() {
  char *p = "zab\0";
  printf("%d", findSubstringInWraproundString(p));
}
// @lc code=end
