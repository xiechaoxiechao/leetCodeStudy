/*
 * @lc app=leetcode.cn id=11 lang=javascript
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
/**
 * @param {number[]} height
 * @return {number}
 */
 var maxArea = function(height) {
    var j=height.length-1;
    var res=0;
    for(var i=0;i!=j;){
        if(height[i]>=height[j]){
            re=height[j]*(j-i);
            j--;
        }else{
            re=height[i]*(j-i)
            i++;
        }
        res=re>res?re:res;
    }
    return(res);
};
// @lc code=end

