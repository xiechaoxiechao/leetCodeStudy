using System;
using System.Collections.Generic;
/*
 * @lc app=leetcode.cn id=71 lang=csharp
 *
 * [71] 简化路径
 */

// @lc code=start
public partial class Solutions {
    public string SimplifyPath(string path) {
        var tem=new List<string>(path.Split('/',System.StringSplitOptions.RemoveEmptyEntries));
        tem.RemoveAll((s)=>s==".");
        for(var i=tem.IndexOf("..");i!=-1;i=tem.IndexOf("..")){
            if(i!=0){
                tem.RemoveRange(i-1,2);
            }else{
                tem.RemoveAt(i);
            }
            
        }

        return "/"+string.Join("/",tem);
         

    }
}
// @lc code=end

