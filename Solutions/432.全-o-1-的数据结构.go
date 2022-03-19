/*
 * @lc app=leetcode.cn id=432 lang=golang
 *
 * [432] 全 O(1) 的数据结构
 */

// @lc code=start
package Solutions

type _node struct {
	key  string
	num  int
	prev *_node
	next *_node
}
type AllOne struct {
	head *_node
	last *_node
	mp   map[string]*_node
}

func Constructor_() AllOne {
	var tem = new(_node)
	return AllOne{tem, tem, make(map[string]*_node)}
}

func (this *AllOne) Inc(key string) {
	if this.head.next == nil {
		tem := _node{key, 1, this.head, nil}
		this.head.next = &tem
		this.last = &tem
		this.mp[key] = &tem
		return
	}
	node, ok := this.mp[key]
	if !ok {
		tem := _node{key, 1, this.last, nil}
		this.last.next = &tem
		this.last = &tem
		this.mp[key] = &tem
		return
	}
	node.num++
	var now *_node
	for now = node.prev; now.num < node.num && now != this.head; now = now.prev {
	}
	if now != node.prev {
		this.mp[now.next.key] = node
		this.mp[key] = now.next
		now.next.key, node.key = node.key, now.next.key
		now.next.num, node.num = node.num, now.next.num
	}

}

func (this *AllOne) Dec(key string) {
	node := this.mp[key]
	node.num--
	if node.num == 0 {
		if node != this.last {
			node.key, this.last.key = this.last.key, node.key
			node.num, this.last.num = this.last.num, node.num
			this.mp[node.key] = node
		}
		this.last = this.last.prev
		this.last.next = nil
		delete(this.mp, key)
		return
	}
	var now *_node
	for now = node.next; now != nil && now.num > node.num; now = now.next {

	}
	if now != node.next {
		if now == nil {
			node.key, this.last.key = this.last.key, node.key
			node.num, this.last.num = this.last.num, node.num
			this.mp[node.key] = node
			this.mp[this.last.key] = this.last
		} else {
			node.key, now.prev.key = now.prev.key, node.key
			node.num, now.prev.num = now.prev.num, node.num
			this.mp[node.key] = node
			this.mp[now.prev.key] = now.prev
		}
	}

}

func (this *AllOne) GetMaxKey() string {
	if this.head.next != nil {
		return this.head.next.key
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	return this.last.key
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
// @lc code=end
