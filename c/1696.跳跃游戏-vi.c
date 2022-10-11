/*
 * @lc app=leetcode.cn id=1696 lang=c
 *
 * [1696] 跳跃游戏 VI
 */

// @lc code=start

#include <stdlib.h>
int findMax(int l, int r, int *nums);
int maxResult(int *nums, int numsSize, int k) {
  int *max = calloc(numsSize, sizeof(int));
  max[0] = nums[0];
  int maxIndex = 0;
  for (int i = 1; i < numsSize; i++) {
    if (maxIndex < i - k) {
      maxIndex = findMax(i - k > 0 ? i - k : 0, i - 1, max);
    }
    if (nums[i] > 0) {
      max[i] = max[maxIndex] + nums[i];
      maxIndex = i;
    } else {
      max[i] = max[maxIndex] + nums[i];
    }
  }
  return max[numsSize - 1];
}
int findMax(int l, int r, int *nums) {
  int maxI = l;
  int max = nums[l];
  for (int i = l + 1; i < r; i++) {
    if (max <= nums[i]) {
      max = nums[i];
      maxI = i;
    }
  }
  return maxI;
}
// @lc code=end
