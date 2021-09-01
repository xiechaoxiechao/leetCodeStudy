/*
 * @lc app=leetcode.cn id=5 lang=javascript
 *
 * [5] 最长回文子串
 */

// @lc code=start
/**
 * @param {string} s
 * @return {string}
 */
 var longestPalindrome = function (s) {
    var l = s.length;
    switch (l) {
        case 0:
            return ''
            break;
        case 1:
            return s
            break;
        case 2:
            if (s[0] == s[1]) {
                return s;
            } else {
                return s[0];
            }
            break;
        default:
            break;
    }
    var a='';
    var c=0;
    for (var i = 1; i < l; i++) {
        var left=i-1,right=i+1;
        for(;s[left]==s[i];left--){}
        for(;s[right]==s[i];right++){}
        for (; left >= 0 && right < l;) {
           if(s[left]==s[right]){
               left--;
               right++;
           }else{
               break;
           }
        }
        var p=right-left-1;
        if(p>c){
            c=p;
            a=s.slice(left+1,right);

        }
        
    }
    return a;

};
// @lc code=end

