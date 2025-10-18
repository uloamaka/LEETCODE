type CustomStack struct {
    stack []int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{ stack: make([]int, 0, maxSize)}
}


func (this *CustomStack) Push(x int)  {
    if len(this.stack) < cap(this.stack) {
        this.stack = append(this.stack, x)
    }  
}


func (this *CustomStack) Pop() int {
    if len(this.stack) ==  0 {
        return -1
    }
    elem := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    return elem
}


func (this *CustomStack) Increment(k int, val int)  {
    m := len(this.stack)
    for i := 0; i < k; i++ {
        if m > 0 {
            this.stack[i] = this.stack[i] + val
            m--
        } else {
            break
        }
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */