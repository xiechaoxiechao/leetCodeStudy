/*
 * @lc app=leetcode.cn id=9 lang=javascript
 *
 * [9] 回文数
 */

// @lc code=start
/**
 * @param {number} x
 * @return {boolean}
 */
 var isPalindrome = function(x) {
    var x=x.toString();
    var n=x.length
   
    for(var i=0;i<=(Math.floor(n/2)+1);i++){
    if(x[i]!==x[n-1-i])
    return false;
    }
    
    return true;

};
// @lc code=end

