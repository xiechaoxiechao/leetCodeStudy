using System;
/*
 * @lc app=leetcode.cn id=64 lang=csharp
 *
 * [64] 最小路径和
 */

// @lc code=start
public partial class Solutions {
    public int MinPathSum(int[][] grid) {
        var stepSum=new int[grid.Length,grid[0].Length];
        stepSum[0,0]=grid[0][0];
        for (int i = 1; i < grid.Length; i++)
        {
            stepSum[i,0]=grid[i][0]+stepSum[i-1,0];
        }
        for (int i = 1; i < grid[0].Length; i++)
        {
            stepSum[0,i]=grid[0][i]+stepSum[0,i-1];
        }
        for (int i = 1; i < grid.Length; i++)
        {
            for (int j = 1; j < grid[0].Length; j++)
            {
                stepSum[i,j]=stepSum[i-1,j]<stepSum[i,j-1]?stepSum[i-1,j]:stepSum[i,j-1];
                stepSum[i,j]+=grid[i][j];
            }
        }
        return stepSum[grid.Length-1,grid[0].Length-1];


    }
}
// @lc code=end

