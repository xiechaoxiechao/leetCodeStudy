using LeetcodeStudy.Solutions.LocationAttrbutes;

namespace LeetcodeStudy.Solutions
{
    [Location("https://leetcode-cn.com/problems/wildcard-matching/")]
    [Description("给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。")]
    public class Solution44
    {

        public bool IsMatch(string s, string p) {
            var isMatch = new bool[s.Length+1,p.Length+1];
            isMatch[0,0] = true;
            for (int i = 1; i <= p.Length; i++)
            {
                if (p[i - 1] == '*')
                {
                    isMatch[0,i] = true;
                }
                else
                {
                    break;
                }
            }
            for (var i = 1; i <= s.Length; i++)
            {
                for (int j = 1; j <= p.Length; j++)
                {
                    if (p[j-1]=='*')
                    {
                        isMatch[i,j] = isMatch[i - 1,j] || isMatch[i,j - 1];
                    }else if (p[j-1] == '?' || p[j-1] == s[i-1])
                    {
                        isMatch[i,j] = isMatch[i - 1,j - 1];
                    }
                }
            }
            return isMatch[s.Length,p.Length];
        }
    }
}