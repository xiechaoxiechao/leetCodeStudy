package Solutions

func countCollisions(directions string) int {
	dire := []byte(directions)
	ans := 0
	for i := 0; i < len(dire); i++ {
		if dire[i] == 'R' {
		} else if dire[i] == 'L' {
			if i > 0 {
				if dire[i-1] == 'R' {
					dire[i-1] = 'S'
					dire[i] = 'S'
					ans += 2
					for j := i - 2; j >= 0; j-- {
						if dire[j] == 'R' {
							dire[j] = 'S'
							ans++

						} else {
							break
						}
					}
				} else if dire[i-1] == 'S' {
					dire[i] = 'S'
					ans++
				}

			}
		} else {
			for j := i - 1; j >= 0; j-- {
				if dire[j] == 'R' {
					ans++
					dire[j] = 'S'
				} else {
					break
				}
			}
		}
	}
	return ans

}
