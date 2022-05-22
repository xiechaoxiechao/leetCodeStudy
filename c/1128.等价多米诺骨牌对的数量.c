/*
 * @lc app=leetcode.cn id=1128 lang=c
 *
 * [1128] 等价多米诺骨牌对的数量
 */

// @lc code=start

#include<stdlib.h>
#include<string.h>
int numEquivDominoPairs(int** dominoes, int dominoesSize, int* dominoesColSize){
    int status[dominoesSize];
    memset(status,0,dominoesSize*sizeof(int));
    int result=0;
    for(int i=0;i<dominoesSize;i++){
        int r=0;
        for(int j=i+1;j<dominoesSize;j++){
            if(status[j]==1||dominoesColSize[i]!=dominoesColSize[j]){
                continue;
            }
            if (memcmp(dominoes[i],dominoes[j],dominoesColSize[i]*sizeof(int))==0){
                status[j]=1;
                r++;
               
            }else{
                int k=0;
                int len=dominoesColSize[i];
                for(;k<len;k++){
                    if(dominoes[i][k]!=dominoes[j][len-k-1]){
                        break;
                    }
                }
                if (k==len){
                    status[j]=1;
                    r++;
                }
            }
            
           
        }
        while (r>0)
        {
            result+=r;
            r--;
        }
        
    }
    return result;

}
// @lc code=end

