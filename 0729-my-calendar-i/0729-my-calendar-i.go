type MyCalendar struct {
	Events [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, interval := range this.Events {
		if start < interval[1] && interval[0] < end {
			return false
		}
	}
	this.Events = append(this.Events, []int{start, end})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */