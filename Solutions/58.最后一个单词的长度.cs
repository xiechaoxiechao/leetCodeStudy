/*
 * @lc app=leetcode.cn id=58 lang=csharp
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
public partial class Solutions {
    public int LengthOfLastWord(string s) {
        var index=s.LastIndexOf(' ');
        if(index==s.Length-1){
            var len=0;
            for(--index;index>=0&&s[index]==' ';--index){};
            for(;index>=0&&s[index]!=' ';--index){
                ++len;
            }
            return len;
        }else{
            return s.Length-1-index;
        }
       

    }
}
// @lc code=end

