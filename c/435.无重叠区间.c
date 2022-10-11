/*
 * @lc app=leetcode.cn id=435 lang=c
 *
 * [435] 无重叠区间
 */

// @lc code=start

#include <stdlib.h>
int compare(int **a, int **b);
int eraseOverlapIntervals(int **intervals, int intervalsSize,
                          int *intervalsColSize) {
  qsort(intervals, intervalsSize, sizeof(intervals[0]), compare);
  int last = intervals[0][1];
  int ans = 0;
  for (int i = 1; i < intervalsSize; i++) {
    if (intervals[i][0] >= last) {
      last = intervals[i][1];
    } else {
      ans++;
    }
  }
  return ans;
}
int compare(int **a, int **b) { return (*a)[1] - (*b)[1]; }
// @lc code=end
