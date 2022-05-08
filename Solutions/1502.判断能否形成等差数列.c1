/*
 * @lc app=leetcode.cn id=1502 lang=c
 *
 * [1502] 判断能否形成等差数列
 */

// @lc code=start

#include<stdbool.h>
#include<stdlib.h>
int cmp(void * a,void * b){
    return *(int*)a-*(int*)b;
}
bool canMakeArithmeticProgression(int* arr, int arrSize){
    qsort(arr,arrSize,sizeof(int),cmp);
    int sub=*(arr+1)-*arr;
    int *end=arr+arrSize;
    for (int *v=arr+2;v<end;v++){
        if(*v-*(v-1)!=sub){
            return false;
        }
    }
    return true;
}
// @lc code=end

