
/*
 * @lc app=leetcode.cn id=75 lang=csharp
 *
 * [75] 颜色分类
 */

// @lc code=start
public partial class Solutions {
    public void SortColors(int[] nums) {
        var left=0;
        var right=nums.Length-1;
        for (int i = 0; i <=right;)
        {
            if(nums[i]==0){
                nums[i]=nums[left];
                nums[left]=0;
                ++left;
                ++i;
                continue;
            }else if(nums[i]==2){
                nums[i]=nums[right];
                nums[right]=2;
                --right;
                continue;
            }
            ++i;
        
        }
    }
    
}
// @lc code=end

