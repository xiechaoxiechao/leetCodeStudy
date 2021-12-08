using System.Collections.Generic;
/*
 * @lc app=leetcode.cn id=48 lang=csharp
 *
 * [48] 旋转图像
 */

// @lc code=start
public partial class Solutions {
    public void Rotate(int[][] matrix) {
        var tem=0;
        var times=matrix.Length/2;
        var ti=matrix.Length-1;
        var tj=matrix.Length-1;
        for(var i=0;i<times;++i){
            for(var j=i;j<matrix.Length-1-i;++j){
               tem=matrix[i][j];
               matrix[i][j]=matrix[tj][i];
               matrix[tj][i]=matrix[ti][tj];
               matrix[ti][tj]=matrix[j][ti];
               matrix[j][ti]=tem;
               --tj;
            }
            --ti;
            tj=ti;

        }
        
    }
}
// @lc code=end

