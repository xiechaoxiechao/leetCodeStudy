using System.Collections;
using System;
using System.Collections.Generic;
/*
 * @lc app=leetcode.cn id=54 lang=csharp
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
public partial class Solution {
    public IList<int> SpiralOrder(int[][] matrix) {
        var res=new int[matrix.Length*matrix[0].Length];
        var length=new int[(matrix.Length<matrix[0].Length?matrix.Length:matrix[0].Length)/2+1];
        if (length.Length>0){
            length[0]=0;
        }
        for(int i=1;i<length.Length;i++){
            length[i]=2*(matrix.Length+matrix[0].Length-2-4*(i-1))+length[i-1];
        }
        var layer=0;
        for(var i=0;i<matrix.Length;++i){
            for(var j=0;j<matrix[i].Length;++j){
                layer=(i<j?i:j)<(matrix.Length-i<matrix[0].Length-j?matrix.Length-i-1:matrix[0].Length-j-1)?(i<j?i:j):(matrix.Length-i-1<matrix[0].Length-j-1?matrix.Length-i-1:matrix[0].Length-j-1);
                if(i==layer||j==matrix[i].Length-1-layer){
                    res[length[layer]+i+j-2*layer]=matrix[i][j];
                }else{
                     res[length[layer+1]-(i+j-2*layer)]=matrix[i][j];
                     
                }
                
            }

        }
        return res;
    }
}
// @lc code=end

