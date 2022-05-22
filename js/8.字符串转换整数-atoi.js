/*
 * @lc app=leetcode.cn id=8 lang=javascript
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
/**
 * @param {string} s
 * @return {number}
 */
 var myAtoi = function(str) {
    var d;
    var c=str.trim();
    if(c.length===0)
    return 0
    if(c[0]!=='-'&&c[0].match(/[0-9]/g)==null&&c[0]!=='+')
    return 0;
    var a=c.match(/([+,-]?[0-9]*)/g);
    for(let i=0,n=a.length;i<n;i++){
      if(a[i]!==''){
          if(a[i]==='-'){
              return 0;
          }
          if(a[i]==='+'){
              return 0;
          }
        d=parseInt(a[i]);
        break;
      }
    }
  
    var b=Math.pow(2,31);
    if(d<-1*b){
        return Math.pow(2,31)*-1;
    }else if(d>b-1){
        return(Math.pow(2,31)-1)
    }else{
        return (d);
    }
  
  };
  
// @lc code=end

