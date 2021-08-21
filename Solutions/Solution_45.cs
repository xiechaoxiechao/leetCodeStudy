using LeetcodeStudy.Solutions.LocationAttrbutes;

namespace LeetcodeStudy.Solutions
{
    [Location("https://leetcode-cn.com/problems/jump-game-ii/")]
    [Description(@"给你一个非负整数数组 nums ，你最初位于数组的第一个位置。

        数组中的每个元素代表你在该位置可以跳跃的最大长度。

        你的目标是使用最少的跳跃次数到达数组的最后一个位置。

        假设你总是可以到达数组的最后一个位置")]
    public class Solution45

    {

        public int Jump(int[] nums)
        {
            if (nums.Length == 1)
            {
                return 0;
            }
            var times = 0;
            var max = 0;
            var max1 = 0;
            for (var i = 0; i < nums.Length; )
            {
                if (i + nums[i] >= nums.Length-1)
                {
                    return times + 1;
                }
                for (var j = 1; j <= nums[i]; ++j)
                {
                    if (max < j + nums[i + j])
                    {
                        max = j + nums[i + j];
                        max1 = j;
                    }


                }
                i += max1;
                max = 0;
                max1 = 0;
                times++;
                if (i > nums.Length-1)
                {
                    return times+1;
                }

            }
            return times;
        }
    }
}