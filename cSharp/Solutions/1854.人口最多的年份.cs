using System;
/*
 * @lc app=leetcode.cn id=1854 lang=csharp
 *
 * [1854] 人口最多的年份
 */

// @lc code=start
public partial class Solution {
    public int MaximumPopulation(int[][] logs) {
        var k=new int[100];
        foreach(var vs in logs){
            for(int i=vs[0];i<vs[1];i++){
                k[i-1950]++;
            }
        }
        var max=Int32.MinValue;
        var maxY=0;
        for(int i=0;i<100;i++){
            if(max<k[i]){
                max=k[i];
                maxY=i;
            }
        }
        return maxY+1950;
    }
}
// @lc code=end

