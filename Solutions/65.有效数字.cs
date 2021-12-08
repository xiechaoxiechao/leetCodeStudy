using System;
using System.Linq;
/*
 * @lc app=leetcode.cn id=65 lang=csharp
 *
 * [65] 有效数字
 */

// @lc code=start
public partial class Solutions
{
    public bool IsNumber(string s)
    {
        s = s.ToLower();
        var eIndex = -1;
        for (var i = 0; i < s.Length; ++i)
        {
            if ('a' <= s[i] && s[i] <= 'z')
            {
                if (s[i] == 'e' && eIndex == -1)
                {
                    eIndex = i;
                    continue;
                }
                return false;
            }
        }
        if (eIndex == 0 || eIndex == s.Length - 1)
        {
            return false;
        }
        char[] num1;
        var lastIndex = -1;
        if (eIndex != -1)
        {
            num1 = new char[eIndex];
            s.CopyTo(0, num1, 0, eIndex);
            var num2 = new char[s.Length - 1 - eIndex];
            s.CopyTo(eIndex + 1, num2, 0, num2.Length);
            lastIndex = -1;
            for (int i = 0; i < num2.Length; ++i)
            {
                if (num2[i] == '-' || num2[i] == '+')
                {
                    lastIndex = i;
                }
            }
            if ((lastIndex != -1 && lastIndex != 0 )|| lastIndex == num2.Length - 1)
            {
                return false;
            }
            if(num2.Contains('.')){
                return false;
            }
        }
        else
        {
            num1 = s.ToCharArray();
        }
        lastIndex=-1;

        for (int i = 0; i < num1.Length; ++i)
        {
            if (num1[i] == '-' || num1[i] == '+')
            {
                lastIndex = i;
            }
        }
        if ((lastIndex != -1 && lastIndex != 0)||lastIndex==num1.Length-1)
        {
            return false;
        }
        var dotIndex = -1;
        for (var i = 0; i < num1.Length; ++i)
        {
                if (num1[i] == '.' )
                {
                    if (dotIndex==-1){
                        dotIndex=i;
                        continue;
                    }
                   return false;
                }
        }
        if (dotIndex == num1.Length - 1&&dotIndex==lastIndex+1)
        {
            return false;
        }












        return true;


    }
}
// @lc code=end

