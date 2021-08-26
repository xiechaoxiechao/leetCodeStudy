using System.Data;
using System;
using System.Linq;
/*
 * @lc app=leetcode.cn id=67 lang=csharp
 *
 * [67] 二进制求和
 */

// @lc code=start
public partial class Solution
{
    public string AddBinary(string a, string b)
    {
        if (a.Length < b.Length)
        {
            var c = a;
            a = b;
            b = c;
        }
        var res = new char[a.Length + 1];
       
        var carry = '0';
        var i = 1;
        for(;i<=b.Length;i++){
             switch (a[a.Length - i] + b[b.Length - i] + carry)
            {
                case 147:
                    res[res.Length - i] = '1';
                    carry = '1';
                    break;
                case 146:
                    res[res.Length - i] = '0';
                    carry = '1';
                    break;
                case 145:
                    res[res.Length - i] = '1';
                    carry = '0';
                    break;
                case 144:
                    res[res.Length-i]='0';
                    break;
            }
        }
       for(;i<=a.Length;i++){
           if(carry!='0'){
               if(a[a.Length-i]=='1'){
                   res[res.Length-i]='0';
               }else{
                   res[res.Length-i]='1';
                   carry='0';
               }
           }else{
               res[res.Length-i]=a[a.Length-i];
           }
       }
        if (carry == '1')
        {
            res[0] = '1';
            return string.Join("", res);

        }
        else
        {
            return string.Join("", res.Skip(1));
        }

    }
}
// @lc code=end

