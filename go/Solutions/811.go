package Solutions

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	var mp = make(map[string]int)
	for i := 0; i < len(cpdomains); i++ {
		spaceIndex := strings.IndexByte(cpdomains[i], ' ')
		count, _ := strconv.Atoi(cpdomains[i][:spaceIndex])
		domain := cpdomains[i][spaceIndex+1:]
		mp[domain] += count
		for strings.Count(domain, ".") > 0 {
			dotIndex := strings.IndexByte(domain, '.')
			domain = domain[dotIndex+1:]
			mp[domain] += count
		}

	}
	var ans = make([]string, 0, len(mp))
	for k, v := range mp {
		ans = append(ans, strconv.Itoa(v)+" "+k)
	}
	return ans
}
