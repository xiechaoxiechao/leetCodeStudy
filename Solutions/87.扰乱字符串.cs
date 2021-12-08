using System.Collections.Generic;

/*
 * @lc app=leetcode.cn id=87 lang=csharp
 *
 * [87] 扰乱字符串
 */

// @lc code=start
public partial class Solutions
{
    private Dictionary<(int, int, int), bool> map = new Dictionary<(int, int, int), bool>();

    public bool IsScramble(string s1, string s2)
    {
        if (s1.Length != s2.Length)
        {
            return false;
        }

        return C(0, 0, s1.Length - 1, s1, s2);
    }

    private bool C(int begin1, int begin2, int len, string s1, string s2)
    {
        var t = (begin1, begin2, len);
        if (map.ContainsKey(t))
        {
            return map[t];
        }
        var end2 = begin2 + len;
        if (len == 0)
        {
            return s1[begin1] == s2[begin2];
        }

        var n1 = new byte[26];
        for (int i = 0; i < len; i++)
        {
            n1[s1[begin1 + i] - 'a'] += 1;
            n1[s2[begin2 + i] - 'a'] -= 1;
        }

        foreach (var i in n1)
        {
            if (i != 0)
            {
                map.Add(t, false);
                return false;
            }
        }


        for (var i = 0; i < len; i++)
        {
            if (C(begin1, end2 - i, i, s1, s2) && C(begin1 + i + 1, begin2, len - i - 1, s1, s2))
            {
                map.Add(t, true);
                return true;
            }

            if (C(begin1, begin2, i, s1, s2) && C(begin1 + i + 1, begin2 + i + 1, len - i - 1, s1, s2))
            {
                map.Add(t, true);
                return true;
            }
        }

        map.Add(t, false);
        return false;
    }
}
// @lc code=end