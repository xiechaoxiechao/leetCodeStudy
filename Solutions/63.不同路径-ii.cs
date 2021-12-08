/*
 * @lc app=leetcode.cn id=63 lang=csharp
 *
 * [63] 不同路径 II
 */

// @lc code=start
public partial class Solutions {
    public int UniquePathsWithObstacles(int[][] obstacleGrid) {
        var step=new int[obstacleGrid.Length,obstacleGrid[0].Length];
        
        for (int i = 0; i < obstacleGrid.Length; ++i)
        {
            if(obstacleGrid[i][0]==1){break;}
            step[i,0]=1;
        }
         for (int i = 0; i < obstacleGrid[0].Length; ++i)
        {
            if(obstacleGrid[0][i]==1){break;}
            step[0,i]=1;
        }
        for (int i = 1; i < obstacleGrid.Length; ++i)
        {
            for (int j = 1; j < obstacleGrid[0].Length; ++j)
            {
                step[i,j]=obstacleGrid[i][j]==1?0:step[i-1,j]+step[i,j-1];
            }
        }
        return step[obstacleGrid.Length-1,obstacleGrid[0].Length-1];

    }
}
// @lc code=end

