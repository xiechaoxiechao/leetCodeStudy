using System.Collections.Generic;

/*
 * @lc app=leetcode.cn id=80 lang=csharp
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
public partial class Solutions
{
    public int RemoveDuplicates(int[] nums)
    {
        var last = nums[0];
        var len = nums.Length;
        var tem = 0;
        var counter = 0;
        var l=0;
        var mov=new List<int>();
        for (var i = 1; i < len;)
        {
            if (last == nums[i])
            {
                counter++;
                if (counter > 1)
                {
                    tem = i;
                    l++;
                    for (i++; i < len && nums[i] == last; i++)
                    {
                        l++;
                    }
                    mov.Add(tem);
                    mov.Add(i);
                    if(i>=len){
                        break;
                    }
                    last=nums[i];
                    i++;
                    counter=0;

                }else{
                    i++;
                    continue;
                }
            }else{
                last=nums[i];
                i++;
                counter=0;
            }
           
        }
        for(var i=mov.Count-1;i>=1;i-=2){
            for(int j=0;j<len-mov[i];j++){
                nums[mov[i-1]+j]=nums[mov[i]+j];
            }
            len-=(mov[i]-mov[i-1]);
        }
        return len;
    }
}
// @lc code=end

