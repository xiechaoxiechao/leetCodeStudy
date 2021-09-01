/*
 * @lc app=leetcode.cn id=3 lang=javascript
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
/**
 * @param {string} s
 * @return {number}
 */
 var lengthOfLongestSubstring = function(s) {
    var a=[];
    var max=0;
    for(var i=0,j=s.length;i<j;i++){
        a.push(s[i])
        for(var m=0,n=a.length;m<n-1;m++){
            if(a[m]===s[i]){
                a.splice(0,m+1);
                break;
            }
        }
        if(max<a.length)
        max=a.length;
    }
    return max;

};
// @lc code=end

