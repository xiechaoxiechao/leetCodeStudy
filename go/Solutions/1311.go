package Solutions

import (
	"fmt"
	"sort"
)

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	var mp1 = map[int]struct{}{}
	sl := friends[id]
	mp1[id] = struct{}{}
	for lev := 1; lev < level; lev++ {
		slT := make([]int, 0, 8)
		for _, v := range sl {
			mp1[v] = struct{}{}
		}

		for i := 0; i < len(sl); i++ {
			for j := 0; j < len(friends[sl[i]]); j++ {

				if _, ok := mp1[friends[sl[i]][j]]; !ok {
					mp1[friends[sl[i]][j]] = struct{}{}
					slT = append(slT, friends[sl[i]][j])
				}
			}

		}
		sl = slT
	}
	fmt.Println(sl)
	var mp2 = map[string]int{}
	for _, v := range sl {
		for _, v1 := range watchedVideos[v] {
			mp2[v1]++
		}
	}
	pairs := make([]pair, 0, len(mp2))
	for k, v := range mp2 {
		pairs = append(pairs, pair{
			k,
			v,
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count < pairs[j].count {
			return true
		} else if pairs[i].count == pairs[j].count {
			return pairs[i].s < pairs[j].s
		} else {
			return false
		}
	})
	ans := make([]string, 0, len(pairs))
	for _, v := range pairs {
		ans = append(ans, v.s)
	}
	return ans

}

type pair struct {
	s     string
	count int
}
