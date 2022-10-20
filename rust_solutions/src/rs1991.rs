/*
 * @lc app=leetcode.cn id=1991 lang=rust
 *
 * [1991] 找到数组的中间位置
 */
use super::solution::Solution;
// @lc code=start
impl Solution {
    pub fn find_middle_index(nums: Vec<i32>) -> i32 {
        let mut pre_sum:Vec<i32>=Vec::with_capacity(nums.len());
        pre_sum.push(0);
        for (i,&v) in nums.iter().enumerate(){
            pre_sum.push(pre_sum[i]+v)  ;
        };
        for i in 0..pre_sum.len()-1{
            if pre_sum[i]==pre_sum.last().unwrap()-pre_sum[i+1]{
                return i as i32
            }
        };
        return  -1;

    }
}
// @lc code=end

