using System;
/*
 * @lc app=leetcode.cn id=66 lang=csharp
 *
 * [66] 加一
 */

// @lc code=start
public partial class Solutions {
    public int[] PlusOne(int[] digits) {
        Action<int> ana=(a)=>{};
        ana=(index)=>{
            if(digits[index]==9){
                digits[index]=0;
                if(index!=0){
                    ana(index-1);
                }else{
                   Array.Resize(ref digits,digits.Length+1);
                   digits[0]=1;
                }
                
            }else{
                ++digits[index];
            }
        };
        ana(digits.Length-1);
        
        return digits;
    }
}
// @lc code=end

