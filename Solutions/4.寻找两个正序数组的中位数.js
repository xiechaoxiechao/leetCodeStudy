/*
 * @lc app=leetcode.cn id=4 lang=javascript
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
 var findMedianSortedArrays = function(nums1, nums2) {
    if(nums1.length===0){
        var k=nums2.length;
        if(k%2){
            return (nums2[(k+1)/2-1]);
        }
        return ((nums2[k/2]+nums2[k/2-1])/2);
    }
    if(nums2.length===0){
        var k=nums1.length;
        if(k%2){
            return (nums1[(k+1)/2-1]);
        }
        return ((nums1[k/2]+nums1[k/2-1])/2);
    }
    var a=[];
    var k1=nums1.length;
    var k2=nums2.length
    for(var n=0,x=0,m=0;;m++){
        if(nums1[n]<nums2[x]){
            a.push(nums1[n]);
            n++
        }else if(nums1[n]==nums2[x]){
            a.push(nums1[n]);
            n++;
            a.push(nums2[x])
            x++;
        }else{
            a.push(nums2[x]);
            x++;
        }
        if(n==k1){
            var res=a.concat(nums2.slice(x,k2));
            var l=res.length;
            if(l%2){
                return(res[(l+1)/2-1]);
            }
            return ((res[l/2]+res[l/2-1])/2)

        }
        if(x==k2){
            var res=a.concat(nums1.slice(n,k1));
            var l=res.length;
            if(l%2){
                return(res[(l+1)/2-1]);
            }
            return ((res[l/2]+res[l/2-1])/2)

        }
    }

};
// @lc code=end

