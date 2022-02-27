using System;
/*
 * @lc app=leetcode.cn id=91 lang=csharp
 *
 * [91] 解码方法
 */

// @lc code=start
using System.Collections.Generic;
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
// @lc code=end

