/*
 * @lc app=leetcode.cn id=462 lang=c
 *
 * [462] 最少移动次数使数组元素相等 II
 */

// @lc code=start

#include<stdio.h>
#include<stdlib.h>
int cmp(void * a,void* b);
int minMoves2(int* nums, int numsSize){
    qsort(nums,numsSize,sizeof(int),cmp);
    long step=0;
    for(int i=1;i<numsSize;i++){
        step+=nums[i]-nums[0];
    }
    int left=1;
    int right=numsSize-1;
    for(int i=1;i<numsSize;i++){
        int a=(nums[i]-nums[i-1])*(left-right);
        if (a>0){
            return step;
        }else{
            step+=a;
            left++;
            right--;
        }
    }
    return step;
}
int cmp(void * a,void* b){
    return *(int*)a-*(int*)b;
}
// @lc code=end

