/*
 * @lc app=leetcode.cn id=2027 lang=csharp
 *
 * [2027] 转换字符串的最少操作次数
 */

// @lc code=start
public partial class Solution {
    public  int MinimumMoves(string s) {
        var res=0;
        for(var i=0;i<s.Length;){
            if(s[i]=='X'){
                res++;
                i+=3;
            }else{
                i++;
            }
        }
        return res;
    }
}
// @lc code=end

