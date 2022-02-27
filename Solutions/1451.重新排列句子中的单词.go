/*
 * @lc app=leetcode.cn id=1451 lang=golang
 *
 * [1451] 重新排列句子中的单词
 */
package Solutions

import (
	"sort"
	"strings"
)

// @lc code=start
func arrangeWords(text string) string {
	strs := strings.Split(text, " ")
	var mp = make(map[int][]string)
	tem := []byte(strs[0])
	tem[0] += 32
	strs[0] = string(tem)
	for i := 0; i < len(strs); i++ {
		tem, ok := mp[len(strs[i])]
		if !ok {
			tem = []string{strs[i]}
		} else {
			tem = append(tem, strs[i])
		}
		mp[len(strs[i])] = tem
	}
	var keys = make([]int, 0, len(mp))
	for key := range mp {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	var sb = strings.Builder{}
	ts := mp[keys[0]]
	sb.WriteByte(ts[0][0] - 32)
	sb.WriteString(ts[0][1:])
	for i := 1; i < len(ts); i++ {
		sb.WriteByte(' ')
		sb.WriteString(ts[i])
	}

	for i := 1; i < len(keys); i++ {
		ts = mp[keys[i]]
		for _, v := range ts {
			sb.WriteByte(' ')
			sb.WriteString(v)
		}
	}
	return sb.String()

}

// @lc code=end
