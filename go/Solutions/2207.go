package Solutions
func maximumSubsequenceCount(text string, pattern string) int64 {
	var pre = make([]int, len(text))
	var max = 0
	num := 0
	for i := range text {
		
		pre[i] = num
        if text[i] == pattern[0] {
			num++
		}
	}
	count := 0
	for i := 0; i < len(text); i++ {
		if text[i] == pattern[1] {
			count += pre[i] + 1
		}
	}
 
	max = count
	num = 0
	for i := len(text) - 1; i >= 0; i-- {
		
		pre[i] = num
        if text[i] == pattern[1] {
			num++
		}
	}
	count = 0
	for i := 0; i < len(text); i++ {
		if text[i] == pattern[0] {
			count += pre[i] + 1
		}
	}
	if max > count {
		return int64(max)
	}
	return int64(count)
}
