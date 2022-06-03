using System;
/*
 * @lc app=leetcode.cn id=961 lang=csharp
 *
 * [961] 在长度 2N 的数组中找出重复 N 次的元素
 */

// @lc code=start
using System.Collections.Generic;
public partial class Solution {
    public int RepeatedNTimes(int[] nums) {
        var map=new HashSet<int>();
        foreach(var v in nums){
            if(map.Contains(v)){
                return v;
            }
            map.Add(v);
        }
        return 0;
    }
}
// @lc code=end

