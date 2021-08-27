/*
 * @lc app=leetcode.cn id=7 lang=javascript
 *
 * [7] 整数反转
 */

// @lc code=start
/**
 * @param {number} x
 * @return {number}
 */
 var reverse = function(x) {
   
    var c=x.toString().split('');
    var n=c.length;
    var d;
    if(c[0]==='-'){
        d=parseInt('-'+c.slice(1,n+1).sort((a,b)=>1).toString().replace(/,/g,''))
    }else{
     d=parseInt(c.sort((a,b)=>1).toString().replace(/,/g,''))
    }
     var b=Math.pow(2,31);
    if(d<-1*b||d>b-1)
    return 0
    return d;
  
  };
// @lc code=end

