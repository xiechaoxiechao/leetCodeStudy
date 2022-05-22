/*
 * @lc app=leetcode.cn id=1416 lang=c
 *
 * [1416] 恢复数组
 */

// @lc code=start

#include <stdlib.h>

int numberOfArrays(char * s, int k){
    char * ks=calloc(11,sizeof(char)) ;
    sprintf(ks,"%d",k);
    int n,n1;
    n1=(int)strlen(ks);
    n=(int )strlen(s);
    int  dp[n+1];
    memset(dp,0,sizeof(int)*(n+1) );
    dp[0]=1;
    for(int i=0;i<n;i++){
        if(s[i]=='0'){
            continue;
        }
        long long nu=0;
        for(int j=0;j<n1&&j+i<n;j++){
            nu*=10;
            nu+=s[i+j]-'0';
            if(nu>(long long)k){
                break;
            }else{
                dp[i+j+1]=(dp[i+j+1]+dp[i])%((int)(1E9+7));
            }
        }
    }
    return dp[n];

}
// @lc code=end

