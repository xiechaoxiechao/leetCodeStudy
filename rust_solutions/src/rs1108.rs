/*
 * @lc app=leetcode.cn id=1108 lang=rust
 *
 * [1108] IP 地址无效化
 */
pub struct Solution{}
// @lc code=start
impl Solution {
    pub fn defang_i_paddr(address: String) -> String {
        address.replace(".", "[.]")
    }
}
// @lc code=end
#[cfg(test)]
mod tests{
    use super::*;
    #[test]
    fn t(){
        assert_eq!(Solution::defang_i_paddr("1.1.1.1".to_string()),"1[.]1[.]1[.]1")
    }
}