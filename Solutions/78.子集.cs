using System;
/*
 * @lc app=leetcode.cn id=78 lang=csharp
 *
 * [78] 子集
 */

// @lc code=start
using System.Collections.Generic;
public partial class Solution {
    public IList<IList<int>> Subsets(int[] nums) {
        var times=1<<nums.Length;
        var res=new List<IList<int>>();
        List<int> tem;
        for(var i =0;i<times;++i){
            tem=new List<int>();
            for(int j=0;(1<<j)<=i;++j){
                if(((1<<j)&i)!=0){
                    tem.Add(nums[j]);
                }
            }
            res.Add(tem);
        }

        return res;
    }
}
// @lc code=end

