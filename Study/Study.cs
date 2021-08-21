using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace LeetcodeStudy.Study
{
    public class Study
    {
        public delegate void Handler();
        public Study()
        {
            Eventing += ana;
            Eventing += ana1;
        }
        /*
         * test
         */
        public void Run()
        {
            OnEvent();
            var a = new int[6] { 1, 2, 3, 4, 5, 6 };
            IEnumerable<int> b = from i in a
                where i < 5
                select i;
            foreach (var i in b)
            {
                Console.Write(i);
            }

            var c = new int[] { 2, 1, 3 };
           test(c);
           Console.WriteLine(c[0]);





        }
        public void test(int[] i){
            i[0]=10;
        }

        public virtual void OnEvent()
        {
            Eventing?.Invoke();
            Console.WriteLine("onEvent");
        }

        public event Handler Eventing;

        public async void ana()
        {
            await Task.Delay(1000);
            Console.WriteLine("ana");
        }
        public  void ana1()
        {
            Console.WriteLine("ana1");
        }
    }
}