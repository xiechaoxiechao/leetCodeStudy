using System.Collections.Generic;
/*
 * @lc app=leetcode.cn id=60 lang=csharp
 *
 * [60] 排列序列
 */

// @lc code=start
public partial class Solution
{
    public string GetPermutation(int n, int k)
    {
        var step = new int[n];
        var res = string.Empty;
        var tem = new List<int>();
        step[0] = 1;
        tem.Add(1);
        for (int i = 1; i < n; i++)
        {
            tem.Add(i + 1);
            step[i] = step[i - 1] * i;
        }
        int aim = 0;
        for (int i = 0; i < n; ++i)
        {
            aim = k / step[n - 1 - i];
            k -= aim * step[n - 1 - i];
            if(k==0){
                --aim;
                res += tem[aim].ToString();
                tem.RemoveAt(aim);
                break;
            }else{
                res += tem[aim].ToString();
                tem.RemoveAt(aim);
            }

        };
        for (int i = tem.Count-1; i >=0; --i)
        {
            res+=tem[i].ToString();
        }

        return res;

    }
}
// @lc code=end

