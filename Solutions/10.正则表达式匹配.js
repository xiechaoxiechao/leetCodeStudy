/*
 * @lc app=leetcode.cn id=10 lang=javascript
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
/**
 * @param {string} s
 * @param {string} p
 * @return {boolean}
 */
 var isMatch = function(s, p) {
    var c=new RegExp(`${p}`);
    var b=s.match(c);
    if(b===null){
      return false;
    }else{
      if(b[0]===s){
        return true;
      }else{
        return false;
      }
    }
  };
// @lc code=end

