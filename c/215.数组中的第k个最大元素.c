/*
 * @lc app=leetcode.cn id=215 lang=c
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start


void verifyHeap(int * heap,int len,int node){
    int max=0;
    int left=node*2+1;
    int right=left+1;
    if (left<len){
        max=heap[node]>heap[left]?0:1;
    }else{
        return;
    }
    if(right<len){
        if (max==0){
            max=heap[node]>heap[right]?0:2;
        }else{
            max=heap[left]>heap[right]?1:2;
        }
    }
    if(max==0){
        return;
    }
    int t=heap[node];
    if (max==1){
        heap[node]=heap[left];
        heap[left]=t;
        verifyHeap(heap,len,left);
        return;
    }
    if (max==2){
        heap[node]=heap[right];
        heap[right]=t;
        verifyHeap(heap,len,right);
    }
}
int findKthLargest(int* nums, int numsSize, int k){
    for(int i=(numsSize-1)/2;i>=0;i--){
        verifyHeap(nums,numsSize,i);
    }
    for(int i=1;i<k;i++){
        int t=nums[0];
        nums[0]=nums[numsSize-i];
        nums[numsSize-i]=t;
        verifyHeap(nums,numsSize-i,0);
    }
    return nums[0];
}
// @lc code=end

