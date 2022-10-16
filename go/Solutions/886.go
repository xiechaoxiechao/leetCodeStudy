package Solutions

type Uf struct {
	master []int
	rank   []int
}

func (u *Uf) find(i int) int {
	if u.master[i] != i {
		u.master[i] = u.find(u.master[i])
	}
	return u.master[i]
}

func (u *Uf) Union(i, j int) {
	x, y := u.find(i), u.find(j)
	if x == y {
		return
	}
	if u.rank[x] > u.rank[y] {
		u.master[y] = x
	} else if u.rank[x] == u.rank[y] {
		u.master[x] = y
		u.rank[y]++
	} else {
		u.master[x] = y
	}
}

func (u *Uf) isConnected(i, j int) bool {
	x, y := u.find(i), u.find(j)
	return x == y
}
func possibleBipartition(n int, dislikes [][]int) bool {
    if len(dislikes)==0{
        return true
    }
	uf := &Uf{}
	uf.master = make([]int, n)
	for i := 0; i < n; i++ {
		uf.master[i] = i
	}
	uf.rank = make([]int, n)
	con := make([][]int, n)
	for i := 0; i < len(dislikes); i++ {
		a, b := dislikes[i][0]-1, dislikes[i][1]-1
		con[a] = append(con[a], b)
		con[b] = append(con[b], a)
	}
	for i := 0; i < len(con); i++ {
		for j := 0; j < len(con[i]); j++ {
			uf.Union(con[i][0], con[i][j])
		}
		if len(con[i])>0&&uf.isConnected(con[i][0], i) {
			return false
		}
	}
	return true

}
