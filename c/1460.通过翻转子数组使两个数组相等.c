/*
 * @lc app=leetcode.cn id=1460 lang=c
 *
 * [1460] 通过翻转子数组使两个数组相等
 */

// @lc code=start
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

int compare(void *i, void *j);
bool canBeEqual(int *target, int targetSize, int *arr, int arrSize) {
  qsort(target, targetSize, sizeof(int), compare);
  qsort(arr, arrSize, sizeof(int), compare);
  return memcmp(target, arr, targetSize * sizeof(int)) == 0;
}
int compare(void *i, void *j) { return *(int *)i - *(int *)j; }
// @lc code=end
