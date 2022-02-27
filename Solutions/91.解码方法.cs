using System;
/*
 * @lc app=leetcode.cn id=91 lang=csharp
 *
 * [91] 解码方法
 */

// @lc code=start
using System.Collections.Generic;
public partial class Solutions {
    public int NumDecodings(string s) {
       var res=1;
       var n=0;
       for(var i=0;i<s.Length;i++){
           if(s[i]>'7'){
               if(n!=0){
                   res*=n;
                   n=0;
               }
           }else if(s[i]>'6'){
               if(n!=0){
                   if(s[i-1]=='1'){
                       if(n==1){
                           res*=2;
                       }else{
                           res*=(2*n-1);
                       }
                       n=0;
                   }
               }
           }else if(s[i]>'2'){
               if(n!=0){
                   if(n==1){
                       res*=2;
                   }else{
                       res*=(2*n-1);
                   }
                   n=0;
               }
           }else if(s[i]>'0'){
               n++;
           }else{
               if(n==0){
                   return 0;
               }else{
                   n--;
               }
           }
       }
       return res;
    }
}
// @lc code=end

