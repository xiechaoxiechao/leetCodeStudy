using System;

using System.Collections.Generic;


namespace LeetcodeStudy.Solutions
{
    
    class Solution47
    {
        private List<IList<int>> _res;
        private Stack<int> _tem;
        private int[] _status;
        private Dictionary<int, int> _re;

        public Solution47()
        {
            _res = new List<IList<int>>();
            _tem = new Stack<int>();
            _re = new Dictionary<int, int>();
        }

        public IList<IList<int>> PermuteUnique(int[] nums) {
            if (nums.Length == 0)
            {
                _res.Add(Array.Empty<int>());
                return _res;
            }

            if (nums.Length == 1)
            {
                _res.Add(new int[]{nums[0]});
                return _res;
            }
            _status = new int[nums.Length];
            Array.Sort(nums);
            var tem = nums[0];
            var times = 1;
            for (var i=1;i<nums.Length;++i)
            {
                if (tem == nums[i])
                {
                    times++;
                }
                else
                {
                    if (times > 1)
                    {
                        _re.Add(tem,0);
                    }
                    tem = nums[i];

                }
            }

            if (times > 1)
            {
                _re.Add(tem,0);
            }

            var di = new Dictionary<int, int>(_re);
            for (int i = 0; i < nums.Length; ++i)
            {
                if (di.ContainsKey(nums[i]))
                {
                    if (di[nums[i]] == 1)
                    {
                        continue;
                    }
                    di[nums[i]] = 1;
                }
                _tem.Push(nums[i]);
                _status[i] = 1;
                Do(ref nums);
                _tem.Pop();
                _status[i] = 0;
            }


            return _res;

        }

        private void Do(ref int[] nums)
        {
            var di = new Dictionary<int, int>(_re);
            for (int i = 0; i < nums.Length; i++)
            {
                if (_status[i] == 1)
                {
                    continue;
                }
                if (di.ContainsKey(nums[i]))
                {
                    if (di[nums[i]] == 1)
                    {
                        continue;
                    }
                    di[nums[i]] = 1;
                }
                _tem.Push(nums[i]);
                _status[i] = 1;
                if (_tem.Count == nums.Length)
                {
                    var n1 = new int[nums.Length];
                    _tem.CopyTo(n1,0);
                    Array.Reverse(n1);
                    _res.Add(n1);
                }
                else
                {
                    Do(ref nums);
                }

                _tem.Pop();
                _status[i] = 0;

            }
        }

    }
}