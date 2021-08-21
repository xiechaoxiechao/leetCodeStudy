using System;
using System.Collections.Generic;
using LeetcodeStudy.Solutions.LocationAttrbutes;

namespace LeetcodeStudy.Solutions
{
    [Location("https://leetcode-cn.com/problems/permutations/")]
    [Description("给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。")]
    public class Solution46
    {
        private IList<IList<int>> _res;
        private int[] _status;
        private List<int> _tem;

        public Solution46()
        {
            _res = new List<IList<int>>();
            _tem = new List<int>();
        }
        public IList<IList<int>> Permute(int[] nums)
        {
            if (nums.Length == 1)
            {
                _res.Add(new int[]{nums[0]});
            }
            _status = new int[nums.Length];
            for (int i = 0; i < nums.Length; i++)
            {
                _status[i] = 1;
                _tem.Add(nums[i]);
                so(ref nums);
                _status[i] = 0;
                _tem.Remove(nums[i]);
            }

            return _res;
        }

        private void so(ref int[] nums)
        {
            for (int i = 0; i < nums.Length; i++)
            {
                if (_status[i] == 1)
                {
                    continue;
                }
                _status[i] = 1;
                _tem.Add(nums[i]);
                if (_tem.Count == nums.Length)
                {
                    var n = new int[nums.Length];
                    _tem.CopyTo(n);
                    _res.Add(n);
                }
                else
                {
                    so(ref nums);

                }
                _status[i] = 0;
                _tem.Remove(nums[i]);

            }
        }
    }
}