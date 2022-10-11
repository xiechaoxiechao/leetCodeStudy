#include <stdbool.h>
#include <stdio.h>
/*
 * @lc app=leetcode.cn id=713 lang=c
 *
 * [713] 乘积小于K的子数组
 */

// @lc code=start
int numSubarrayProductLessThanK(int *nums, int numsSize, int k) {
  if (k == 0 || k == 1) {
    return 0;
  }
  int *left = nums;
  int *right = nums;
  int mul = 1;
  int result = 0;
  nums += numsSize;
  while (right < nums) {
    mul *= *right;
    if (mul < k) {
      right++;
      if (right == nums) {
        int t = right - left;
        printf("%d", t);
        for (int t = right - left; t > 0; t--) {
          result += t;
        }
        return result;
      }
    } else {
      result += (right - left);
      mul /= *right;
      mul /= *left;
      left++;
      if (left > right) {
        mul = 1;
        right++;
      }
    }
  }
  return result;
}

// @lc code=end
