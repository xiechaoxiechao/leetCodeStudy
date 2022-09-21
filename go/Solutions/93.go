package Solutions

func restoreIpAddresses(s string) []string {
	var ans = make([]string, 0, 8)
	track := [4]string{}
	backTrack(s, 0, 3, track, &ans)
	return ans
}

func backTrack(s string, index int, dotNum int, track [4]string, ans *[]string) {
	if dotNum == 0 {
		if s[index] != '0' || index == len(s)-1 {
			if n, _ := strconv.Atoi(s[index:]); n < 256 {
				track[3] = s[index:]
				*ans = append(*ans, strings.Join(track[:], "."))
			}

		}
		return
	}
	if s[index] == '0' {
		track[3-dotNum] = "0"
		if index+1<len(s){
			backTrack(s, index+1, dotNum-1, track, ans)
		}
		
		return
	}
	for i := index; i < len(s)-1; i++ {
		if n, _ := strconv.Atoi(s[index : i+1]); n < 256 {
			track[3-dotNum] = s[index : i+1]
			backTrack(s, i+1, dotNum-1, track, ans)
		}

	}

}
