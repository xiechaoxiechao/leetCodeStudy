using System;
using System.Collections.Generic;
/*
 * @lc app=leetcode.cn id=57 lang=csharp
 *
 * [57] 插入区间
 */

// @lc code=start
public partial class Solutions {
    public int[][] Insert(int[][] intervals, int[] newInterval) {
        var res=new  List<int[]>();
        var end=-1;
        var i=0;
        for(;i<intervals.Length;++i){
            if(newInterval[0]<=intervals[i][0]){
                break;
            }else if(newInterval[0]<=intervals[i][1]){
                newInterval[0]=intervals[i][0];
                break;
            }else{
                res.Add(intervals[i]);
            }
        }
        if (i ==intervals.Length){
            res.Add(newInterval);
            return res.ToArray();
        }
        for(;i<intervals.Length;++i){
             if(newInterval[1]<intervals[i][0]){
                 end=newInterval[1];
                res.Add(new int[]{newInterval[0],end});
                break;
            }else if(newInterval[1]<=intervals[i][1]){
                end=intervals[i][1];
                 res.Add(new int[]{newInterval[0],end});
                 ++i;
                break;
            }
        }
        if (end ==-1){
            res.Add(new int[]{newInterval[0],newInterval[1]});
            return res.ToArray();
        }
        for(;i<intervals.Length;++i){
            res.Add(intervals[i]);
        }
        return res.ToArray();

    }
}
// @lc code=end

