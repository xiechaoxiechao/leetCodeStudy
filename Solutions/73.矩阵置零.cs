/*
 * @lc app=leetcode.cn id=73 lang=csharp
 *
 * [73] 矩阵置零
 */

// @lc code=start
public partial class Solutions {
    public void SetZeroes(int[][] matrix) {
        var cl=false;
        for(int i=0;i<matrix.Length;++i){
            if(matrix[i][0]==0){
                    cl=true;
            }
            for(int j=1;j<matrix[0].Length;++j){
                
                if(matrix[i][j]==0){
                    matrix[i][0]=0;
                    matrix[0][j]=0;
                }

            }
        }
        for(int i=matrix.Length-1;i>=0;--i){
            for(int j=1;j<matrix[0].Length;++j){
                if(matrix[i][0]==0||matrix[0][j]==0){
                    matrix[i][j]=0;
                }
            }
            if(cl){
                matrix[i][0]=0;
            }
        }
    }
}
// @lc code=end

