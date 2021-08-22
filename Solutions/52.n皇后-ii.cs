using System;
using System.Collections.Generic;
/*
 * @lc app=leetcode.cn id=52 lang=csharp
 *
 * [52] N皇后 II
 */

// @lc code=start
public partial class Solution {
    public int TotalNQueens(int n) {
        var status = new int[2*n];
        var length=0;
        var result =0;
        var temRes =0;
        Action add = null;
        var sw = false;
        add = (() =>
        {
            for (int i = 0; i < n; ++i)
            {
                for (int j = 0; j < length - 1; j += 2)
                {
                    if (temRes == status[j] || i == status[j + 1] || (Math.Abs(temRes - status[j]) == Math.Abs(i - status[j + 1])))
                    {
                        sw = true;
                        break;
                    }
                }
                if (sw == true)
                {
                    sw = false;
                    continue;
                }
                status[length]=temRes;
                status[length+1]=i;
                length+=2;
                ++temRes;
                if (temRes== n)
                {
                    ++result;
                }
                else
                {
                    add();
                }
                temRes--;
                length-=2;
            }
        });
        add();
        return result;
    }
}
// @lc code=end

