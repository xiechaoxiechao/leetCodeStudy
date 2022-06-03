/*
 * @lc app=leetcode.cn id=263 lang=c
 *
 * [263] 丑数
 */

// @lc code=start

#include<stdbool.h>
bool isUgly(int n){
    // if (n==1){
    //     return true;
    // };
    if (n<1){
        return false;
    }
    while(n%2==0){
        n/=2;
    };
    while (n%3==0){
        n/=3;
    };
    while(n%5==0){
        n/=5;
    };
    return n==1;
} 
// @lc code=end

