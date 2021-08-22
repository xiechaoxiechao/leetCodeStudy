/*
 * @lc app=leetcode.cn id=53 lang=csharp
 *
 * [53] 最大子序和
 */

// @lc code=start
public partial class Solution
{
    public int MaxSubArray(int[] nums)
    {
        var max = nums[0];
        var tem = 0;
        var tem1 = max;
        var flag=tem1>0?-1:1;
        for (int i = 1; i < nums.Length;)
        {
            if (flag * nums[i] <= 0)
            {
                max=max<nums[i]?nums[i]:max;
                tem1 += nums[i];
                ++i;
            }
            else
            {
                tem += tem1;
                max = max < tem ? tem : max;
                tem=tem<0?0:tem;
                tem1=nums[i];
                ++i;
                flag *= -1;
            }
        }
       if (tem1<0){
           return max;
       }
        max = max < tem+tem1 ? tem+tem1 : max;
        return max;
    }
}
// @lc code=end

