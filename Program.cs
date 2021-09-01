
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
            // var s = new Study.Study();
            // s.Run();
            var a=new string[]{"Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"};
            // var b=new int[2]{2,5};
            // a[0]=new int[]{1,3,1};
            // a[1]=new int[]{1,5,1};
            // a[2]=new int[]{4,2,1};
            // a[3]=new int[]{15,18};
            // a[4]=new int[]{21,22,23,24,25};
            var c=new Solution();
            // ListNode last=null;
            // ListNode tem=null;
            // for(int i=3;i!=0;i--){
            //     tem=new ListNode(i,last);
            //     last=tem;
            // }
           
            // var tes=c.RotateRight(tem,2000000000);
            // for(var i=tes;i!=null;i=i.next){
            //     Console.WriteLine(i.val);
            // }
            Console.WriteLine(c.SimplifyPath("/a/./b/../../c/"));
            // Console.WriteLine(6.ToString());
            
            
            
        }

    }
}
