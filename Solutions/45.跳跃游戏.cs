
namespace LeetcodeStudy.Solutions
{
   
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