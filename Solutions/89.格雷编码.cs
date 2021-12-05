/*
 * @lc app=leetcode.cn id=89 lang=csharp
 *
 * [89] 格雷编码
 */

// @lc code=start
using System.Collections.Generic;
public partial class Solution {
    public IList<int> GrayCode(int n) {
        var res =new int[1<<n];
        res[0]=0;
        res[1]=1;
        var tile=1;
        var head=1;
        for(int i=1;i<n;i++){
            for(int j=tile;j>-1;j--){
                head++;
                res[head]=res[j]+(1<<i);
            }
            tile=head;
        }   
        return res;

    }
}
// @lc code=end

