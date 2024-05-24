type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{queue1: []int{}, queue2: []int{}}
}

func (this *MyStack) Push(x int) {
	this.queue2 = append(this.queue2, x)
	for len(this.queue1) > 0 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	this.queue1, this.queue2 = this.queue2, this.queue1
}

func (this *MyStack) Pop() int {

	if this.Empty() {
		panic("pop from the empty stack")
	}
	topElement := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return topElement
}

func (this *MyStack) Top() int {
	if this.Empty() {
		panic("Top from the empty stack")
	}
	return this.queue1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */