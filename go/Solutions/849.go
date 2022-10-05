package Solutions

func maxDistToClosest(seats []int) int {
	last := 0
	if seats[0] != 1 {
		last = -1
	}
	var max = -1
   
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			if last == -1 {
				max = i 
			} else {
				t := (i - last - 1) / 2
                if (i-last-1)%2==0{
                    t = (i - last - 1) / 2
                }else{
                     t = (i - last ) / 2
                }
				if t > max {
					max = t
				}
			}
			last = i
		}
	}
	if seats[len(seats)-1] != 1 {
		if max < len(seats)-last-1 {
			 max = len(seats)-last-1
		}
	}
	return max
}