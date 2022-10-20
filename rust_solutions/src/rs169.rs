/*
 * @lc app=leetcode.cn id=169 lang=rust
 *
 * [169] 多数元素
 */
use super::solution::Solution;
// @lc code=start
use std::collections::HashMap;
impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let n=nums.len()as i32/2;
        let mut mp :HashMap<i32,i32>=HashMap::new();
        for v in nums.iter(){
            let count=mp.entry(*v).or_insert(0);
            *count+=1;
            if *count>n{
                return *v ;
            }
        };
        -1
    }
}
// @lc code=end

