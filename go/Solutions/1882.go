package Solutions

import "container/heap"

type server struct {
	weight int
	index  int
}

type _servers []*server

func (s _servers) Len() int {
	return len(s)
}
func (s _servers) Less(i, j int) bool {
	if s[i].weight < s[j].weight {
		return true
	} else if s[i].weight == s[j].weight {
		return s[i].index < s[j].index
	} else {
		return false
	}
}
func (s _servers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s *_servers) Pop() interface{} {
	n := len(*s)
	item := (*s)[n-1]
	(*s)[n-1] = nil
	*s = (*s)[:n-1]
	return item
}
func (s *_servers) Push(value interface{}) {
	v := value.(*server)
	*s = append(*s, v)
}

type lastTime []int

func (l lastTime) Len() int {
	return len(l)
}
func (l lastTime) Less(i, j int) bool {
	return l[i] < l[j]
}
func (l lastTime) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l *lastTime) Push(value interface{}) {
	v := value.(int)
	*l = append(*l, v)
}
func (l *lastTime) Pop() interface{} {
	item := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return item
}

func assignTasks(servers []int, tasks []int) []int {
	var ans = make([]int, len(tasks))
	var priorityQue _servers = make([]*server, 0, len(servers))
	var lt lastTime = make([]int, 0)
	heap.Init(&lt)
	for i, v := range servers {
		priorityQue = append(priorityQue, &server{
			v,
			i,
		})
	}
	heap.Init(&priorityQue)
	var time = 0
	var mp = map[int][]*server{}
	for i := 0; i < len(tasks); {
		if s, ok := mp[time]; ok {
			for i := 0; i < len(s); i++ {
				heap.Push(&priorityQue, s[i])
			}
			delete(mp, time)
		}
		for ; i <= time && i < len(tasks); i++ {
			if len(priorityQue) == 0 {
				break
			} else {
				seT := heap.Pop(&priorityQue)
				se := seT.(*server)
				if v, ok := mp[tasks[i]+time]; !ok {
					mp[tasks[i]+time] = []*server{se}
				} else {
					mp[tasks[i]+time] = append(v, se)
				}
				heap.Push(&lt, tasks[i]+time)
				ans[i] = se.index
			}
		}
		if len(priorityQue) == 0 {
			time = heap.Pop(&lt).(int)
		} else {
			time++
		}

	}
	return ans
}
