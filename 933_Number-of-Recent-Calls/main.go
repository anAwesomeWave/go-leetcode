package main

type RecentCounter struct {
	qLen int
	q    []int // очередь
}

func Constructor() RecentCounter {
	return RecentCounter{
		0,
		make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.qLen += 1
	this.q = append(this.q, t)
	for this.q[0] < t-3000 {
		this.qLen -= 1
		this.q = this.q[1:]
	}
	return this.qLen
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
