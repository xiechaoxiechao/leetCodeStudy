/*
 * @lc app=leetcode.cn id=1716 lang=c
 *
 * [1716] 计算力扣银行的钱
 */

// @lc code=start


int totalMoney(int n){
    int a=n/7;
    int ans=0;
    if (a!=0){
        ans+=a*28+7*((a-1)*(a))/2;
    }
    int a1=n%7;
    ans+=a1*a;
    for(int i=1;i<=a1;i++){
        ans+=i;
    }
    return ans;
    

}
// @lc code=end

