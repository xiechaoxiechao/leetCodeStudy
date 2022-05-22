using System;
/*
 * @lc app=leetcode.cn id=50 lang=csharp
 *
 * [50] Pow(x, n)
 */

// @lc code=start
public partial class Solutions
{
    public double MyPow(double x, int n)
    {
        var tem = x;
        double result = 1;
        int isNegative = 0;
        if (n < 0)
        {
            if (n == int.MinValue)
            {
                n = int.MaxValue;
                isNegative = 2;
            }
            else
            {
                n *= -1;
                isNegative = 1;
            }


        }
        for (; n > 0; n /= 2)
        {
            if (n % 2 == 1)
            {

                result *= tem;

            }
            tem *= tem;

        }

        if (isNegative == 1)
        { 
            return 1 / result; 
        }
        else if (isNegative==2){
            return 1 / (result*x);
        }
        return result;
    }
}
// @lc code=end

