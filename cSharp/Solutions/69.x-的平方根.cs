/*
 * @lc app=leetcode.cn id=69 lang=csharp
 *
 * [69] x 的平方根
 */

// @lc code=start
public partial class Solutions {
    public int MySqrt(int x) {
        double c=x;
        while(c*c-x>=1){
            c=c/2+x/(2*c);
            
        }
        return (int)c;
        

    }
}
// @lc code=end

