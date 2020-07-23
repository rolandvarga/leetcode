package main_test

type RecentCounter struct {
	Items []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.Items = append(this.Items, t)

	for i, it := range this.Items {
		if it >= t-3000 {
			this.Items = this.Items[i:]
			break
		}
	}
	return len(this.Items)
}
