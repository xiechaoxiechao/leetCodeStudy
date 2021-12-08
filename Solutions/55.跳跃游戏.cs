/*
 * @lc app=leetcode.cn id=55 lang=csharp
 *
 * [55] 跳跃游戏
 */

// @lc code=start
public partial class Solutions {
    public bool CanJump(int[] nums) {
        int max;
        int tem;
        if (nums.Length==1){
            return true;
        }
        for(int i=0;i<nums.Length;){
            tem=0;
            max=0;
            if(i+nums[i]>=nums.Length-1){
                return true;
            }
            for (int j = 1; j <= nums[i]; j++)
            {
                if(max<j+nums[j+i]){
                    max=j+nums[j+i];
                    tem=j;
                }
            }
            if (tem==0){
                return false;
            }
            i+=tem;
        }
        return false;
    }
}
// @lc code=end

