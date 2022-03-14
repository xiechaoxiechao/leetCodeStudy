package Solutions

func addBinary(a string, b string) string {
	var carry bool
	var maxS string
	var minS string
	if len(a) < len(b) {
		maxS = b
		minS = a
	} else {
		maxS = a
		minS = b
	}
	var res = make([]byte, len(maxS)+1)
	resI := len(res) - 1
	i1, i2 := len(maxS)-1, len(minS)-1
	for i2 >= 0 {
		if carry {
			if maxS[i1] == '1' {
				if minS[i2] == '1' {
					res[resI] = '1'
				} else {
					res[resI] = '0'
				}
			} else {
				if minS[i2] == '1' {
					res[resI] = '0'
				} else {
					res[resI] = '1'
					carry = false
				}
			}

		} else {
			if maxS[i1] == '1' {
				if minS[i2] == '1' {
					carry = true
					res[resI] = '0'
				} else {
					res[resI] = '1'
				}
			} else {
				if minS[i2] == '1' {
					res[resI] = '1'
				} else {
					res[resI] = '0'
				}
			}
		}

		i1--
		i2--
		resI--
	}
	for ; i1 >= 0; i1-- {
		if carry {
			if maxS[i1] == '1' {
				res[resI] = '0'
			} else {
				carry = false
				res[resI] = '1'
			}
		} else {
			if maxS[i1] == '1' {
				res[resI] = '1'
			} else {
				res[resI] = '0'
			}
		}
		resI--

	}
	if carry {
		res[0] = '1'
	} else {
		res[0] = '0'
	}
	if res[0] == '1' {
		return string(res)
	} else {
		return string(res[1:])
	}

}
