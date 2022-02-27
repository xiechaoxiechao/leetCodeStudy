/*
 * @lc app=leetcode.cn id=287 lang=csharp
 *
 * [287] 寻找重复数
 */

// @lc code=start
public partial class Solution {
    public int FindDuplicate(int[] nums) {
        var sum=0;
        foreach(var i in nums){
            sum+=i;
        }
        return sum-(nums.Length*(nums.Length-1)>>1);
    }
}
// @lc code=end

