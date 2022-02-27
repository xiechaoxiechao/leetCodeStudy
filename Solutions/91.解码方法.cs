using System;
/*
 * @lc app=leetcode.cn id=91 lang=csharp
 *
 * [91] 解码方法
 */

// @lc code=start
using System.Collections.Generic;
<<<<<<< HEAD
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
=======
// public partial class Solutions {
//     public int NumDecodings(string s) {
//         var res=1;
//         var c=new char[s.Length];
//         var cIndex=0;
//         var st=new List<int>(s.Length);
//         if(s[0]=='0'){
//             return 0;
//         }
       
//         int i=0;
//         for(;(s[i]=='1'||s[i]=='2')&&i<s.Length;i++){}
//         c[cIndex]=s[i];
//         for(;i<s.Length;i++){
//             if(s[i]=='1'||s[i]=='2'){
//                 cIndex++;
//                 c[cIndex]='a';
//             }else if('2'<s[i]&&s[i]<'7'&&c[cIndex]=='a'){
//                 cIndex++;
//                 c[cIndex]='b';
//             }
//         }    
//         Func< int,int,int> cu=(l,r)=>{
            
//             return 1;
//         };
//         for(int i=1;i<st.Count;i++){
//             res*=cu(st[i-1],st[i]);
//         }

//         return res;
//     }
// }
>>>>>>> cef880380f53f27ae9e79cef5dde4c6524d3ac1d
// @lc code=end

