/*
 * @lc app=leetcode.cn id=70 lang=csharp
 *
 * [70] 爬楼梯
 */

// @lc code=start
public partial  class Solutions {
    public int ClimbStairs(int n) {
        if (n==1){
            return 1;
        }
        if(n==2){
            return 2;
        }
        var last1=1;
        var last2=2;
        var res=0;
        for(int i=3;i<=n;++i){
            res=last1+last2;
            last1=last2;
            last2=res;
        }
        return res;

    }
}
// @lc code=end

