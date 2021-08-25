using System;
using System.Threading.Tasks;
/*
 * @lc app=leetcode.cn id=62 lang=csharp
 *
 * [62] 不同路径
 */

// @lc code=start
public partial class Solution {
    public int UniquePaths(int m, int n) {
        long x=1;
        long y=1;
        var tem=n+m-2;
        var times=m>n?n:m;
        for(--times;times>0;--times){
            x*=tem--;
            y*=times;
        }
        return (int)(x/y); 
    }
}
// @lc code=end

