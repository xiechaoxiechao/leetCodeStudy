using System;
using System.Collections.Generic;
/*
 * @lc app=leetcode.cn id=56 lang=csharp
 *
 * [56] 合并区间
 */

// @lc code=start
public partial class Solution
{
    public int[][] Merge(int[][] intervals)
    {
        Array.Sort(intervals, (t1, t2) =>
        {
            if (t1[0] < t2[0])
            {
                return -1;
            }
            else if (t1[0] > t2[0])
            {
                return 1;
            }
            else
            {
                return 0;
            }
        });
        var res = new List<int[]>();
        for (int i = 0; i < intervals.Length - 1; ++i)
        {
            if (intervals[i][1] < intervals[i + 1][0])
            {
                res.Add(intervals[i]);
            }
            else if (intervals[i][1] >= intervals[i + 1][1])
            {
                intervals[i + 1] = intervals[i];
            }
            else
            {
                intervals[i + 1][0] = intervals[i][0];
            }
        }
        res.Add(intervals[intervals.Length - 1]);
        return res.ToArray();

    }
}
// @lc code=end

