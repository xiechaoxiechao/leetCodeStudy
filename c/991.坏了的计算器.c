/*
 * @lc app=leetcode.cn id=991 lang=c
 *
 * [991] 坏了的计算器
 */

// @lc code=start

int brokenCalc(int startValue, int target) {
  int times = 0;
  while (startValue < target) {
    startValue *= 2;
    times++;
  }
  int a = times;
  int p = target - startValue;
  while (p != 0) {
    times += p / (1 << a);
    p = p % (1 << a);
    a--;
  }
  return times;
}
// @lc code=end
