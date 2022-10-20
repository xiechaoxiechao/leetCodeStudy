/*
 * @lc app=leetcode.cn id=191 lang=rust
 *
 * [191] 位1的个数
 */
use super::solution::Solution;
// @lc code=start
impl Solution {
    pub fn hammingWeight ( mut n: u32) -> i32 {
        let mut ans=0;
        while n>0{
            if n%2!=0{
                ans+=1;
            }
            n>>=1;
        }
       
        ans
    }
}
// @lc code=end

