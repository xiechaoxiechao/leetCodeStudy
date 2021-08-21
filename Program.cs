using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using LeetcodeStudy.Study;
using LeetcodeStudy.Solutions;

namespace LeetcodeStudy
{
    class Program
    {
        static void Main(string[] args)
        {
            var solution = new Solution47();
            var s = new Study.Study();
            s.Run();
            var a=new int[]{2,2,1,1};
            solution.PermuteUnique(a);
            var c=new Solution();
            
        }

    }
}
