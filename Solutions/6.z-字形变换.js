/*
 * @lc app=leetcode.cn id=6 lang=javascript
 *
 * [6] Z 字形变换
 */

// @lc code=start
/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
 var convert = function(s, numRows) {
    var n=s.length;
    var j=numRows;
    var c=new Array(j);
    var g;
    if(j==1){
      return s;
    }else{
      g=Math.floor((n/(2*j-2))+1)*(j-1);
    }
    if(g<j){
      g=j;
    }
    var i;
    if(n===0){
      return '';
    }
    for(i=0;i<j;i++){
      c[i]=new Array(g);
    }
    var swq=0;
    var x=-1;
    var y=0;
    for(i=0;i<n;i++){
      if(swq==0){
        x++;
        c[x][y]=s[i];
        if(x==j-1)
        swq=1;
      }else{
        x--;
        y++;
        c[x][y]=s[i];
        if(x==0)
        swq=0;
      }
    }
    var count;
    var res1=[];
    for(i=0;i<j;i++){
      for(count=0;count<g;count++){
        if(c[i][count]!=undefined){
          res1.push(c[i][count]);
        }
      }
    }
    var data='';
   for(var i=0,m=res1.length;i<m;i++){
     data+=res1[i];
   }
    return (data);
  };
// @lc code=end

