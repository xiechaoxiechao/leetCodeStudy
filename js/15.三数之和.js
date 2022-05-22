/*
 * @lc app=leetcode.cn id=15 lang=javascript
 *
 * [15] 三数之和
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
 var threeSum = function(nums) {
    nums.sort((a,b)=>{
      if(a>b){
        return 1;
      }else{
        return -1;
      }
    })
    var i;
    var n;
    var mid=nums[0];
    var count=0;
    var newArry=[];
    var res=[];
    var a=2;
    for(i=0,n=nums.length;i<n;i++){
        if(nums[i]===mid){
          count++;
          if(mid==0){
            a=3;
          }else{
            a=2;
          }
          if(count<=a){
            newArry.push(nums[i]);
          }
        }else{
          count=1;
          mid=nums[i];
          newArry.push(nums[i]);
        }
      }
    if(newArry.length===0){
      return [];
    }
    var b;
    var j,k;
  
    for(i=1,n=newArry.length;i<n-1;i++){
       j=0;
       k=n-1;
       if(newArry[i]===newArry[i-1]){
         b=2*newArry[i];
         for(;k!=i;k--){
           if(b===-1*newArry[k]){
             res.push([newArry[i],newArry[i],newArry[k]]);
             break;
           }
         }
       }else{
        for(;j!=i&&k!=i;){
          if(newArry[j]===newArry[j-1]){
            j++;
            continue;
          }
          if(newArry[k]===newArry[k+1]){
           k--;
           continue;
         }
          b=newArry[j]+newArry[i]+newArry[k];
          if(b<0){
            j++;
          }else if(b===0){
            res.push([newArry[j],newArry[i],newArry[k]]);
            j++;
            k--;
          }else{
            k--;
          }
        }
       }
       
    }
    return res;
  
  };
// @lc code=end

