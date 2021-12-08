using System;
using System.Collections.Generic;


public partial class Solutions
{
    public string Multiply(string num1, string num2)
    {
        var res = new List<int>();
        int mulRes, tem;
        var d1 = num1.ToCharArray();
        Array.Reverse(d1);
        var d2 = num2.ToCharArray();
        Array.Reverse(d2);
        for (var i = 0; i < d1.Length; ++i)
        {
            tem = 0;
            var j = 0;
            for (; j < d2.Length; ++j)
            {
                mulRes = (d1[i] - '0') * (d2[j] - '0') + tem;
                tem = mulRes / 10;
                mulRes -= tem * 10;
                if (i + j >= res.Count)
                {
                    res.Add(mulRes);
                }
                else
                {
                    if (res[i + j] + mulRes > 9)
                    {
                        res[i + j] += mulRes - 10;
                        ++tem;
                    }
                    else
                    {
                        res[i + j] += mulRes;
                    }
                }
            }

            if (tem != 0)
            {
                if (i + j >= res.Count)
                {
                    res.Add(tem);
                }
                else
                {
                    res[i + j] += tem;
                }
            }
        }

        res.Reverse();
        while (res.Count != 0)
        {
            if (res[0] == 0)
            {
                res.Remove(0);
            }
            else
            {
                break;
            }
        }

        if (res.Count == 0)
        {
            return "0";
        }

        return String.Join("", res);
    }
}