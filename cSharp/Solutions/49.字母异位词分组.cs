using System.Collections.Generic;
using System;
/*
 * @lc app=leetcode.cn id=49 lang=csharp
 *
 * [49] 字母异位词分组
 */

// @lc code=start
public partial class Solutions {
    public IList<IList<string>> GroupAnagrams(string[] strs) {
            var di=new Dictionary<string,List<string>>();
            char[] tem;
            var res=new List<IList<string>>();
            var k=string.Empty;
            foreach(var i in strs){
                tem=i.ToCharArray();
                Array.Sort(tem);
                k=string.Join("",tem);
                if (di.ContainsKey(k)){
                    di[k].Add(i);
                } else{
                    di.Add(k,new List<string>{i});
                }
            }
            foreach(var key in di.Keys){
               res.Add(di[key]);
            }
            return res;

    }
}
// @lc code=end

