/*
 * @lc app=leetcode.cn id=1 lang=javascript
 *
 * [1] 两数之和
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
 var twoSum = function(nums, target) {
    for(var i=0,a=nums.length;i<a;i++){
        for(var j=i+1;j<a;j++ ){
            if(nums[i]+nums[j]===target){
                return[i,j];
            }
        }
    
    }
};
// @lc code=end

