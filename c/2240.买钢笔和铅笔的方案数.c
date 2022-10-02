/*
 * @lc app=leetcode.cn id=2240 lang=c
 *
 * [2240] 买钢笔和铅笔的方案数
 */

// @lc code=start
long long waysToBuyPensPencils(int total, int cost1, int cost2){
    long long ans=0;
    for(int i=0;i*cost1<=total;i++){
        ans++;
        ans+=(total-cost1*i)/cost2;
    }
    return ans;
}


// @lc code=end
