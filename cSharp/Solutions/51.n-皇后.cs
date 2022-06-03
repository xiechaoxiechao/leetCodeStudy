using System;
using System.Collections.Generic;
/*
 * @lc app=leetcode.cn id=51 lang=csharp
 *
 * [51] N 皇后
 */

// @lc code=start
public partial class Solutions
{
    public IList<IList<string>> SolveNQueens(int n)
    {
        var status = new int[2*n];
        var length=0;
        var result = new List<IList<string>>();
        var temRes = new Stack<string>();
        char[] cp1;
        List<string> cp2;
        Action add = ()=>{};
        var sw = false;
        add = (() =>
        {
            for (int i = 0; i < n; ++i)
            {
                for (int j = 0; j < length-1; j += 2)
                {
                    if (temRes.Count == status[j] || i == status[j + 1] || (Math.Abs(temRes.Count - status[j]) == Math.Abs(i - status[j + 1])))
                    {
                        sw = true;
                        break;
                    }
                }
                if (sw == true)
                {
                    sw = false;
                    continue;
                }
                status[length]=temRes.Count;
                status[length+1]=i;
                length+=2;
                cp1 = new char[n];
                Array.Fill<char>(cp1, '.');
                cp1[i] = 'Q';
                temRes.Push(string.Join("", cp1));
                if (temRes.Count == n)
                {
                    cp2 = new List<string>(temRes);
                    cp2.Reverse();
                    result.Add(cp2);
                }
                else
                {
                    add();
                }

                temRes.Pop();
                length-=2;
            }
        });
        add();
        return result;
    }

}
// @lc code=end    

