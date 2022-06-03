/*
 * @lc app=leetcode.cn id=944 lang=c
 *
 * [944] 删列造序
 */

// @lc code=start

#include<stdio.h>
#include<stdlib.h>
#include<string.h>
int minDeletionSize(char ** strs, int strsSize){
    int len=strlen(*strs);
    if(strsSize==1){
        return 0;
    }
    int res=0;
    for(int i=0;i<len;i++){
        int j=1;
        for (;j<strsSize;j++){
            if(strs[j][i]<strs[j-1][i]){
                break;
            }
        }
        if (j!=strsSize){
            res++;
        }

    }
    return res;

}
int main(){
    char *c[3][4]={"rrjk","furt","guzm"};
    int k=strlen(c[0]);
    minDeletionSize(c,3);
}
// @lc code=end

