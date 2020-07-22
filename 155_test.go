type MinStack struct {
    Elements []int
    Min []int
}

func Constructor() MinStack {
    return MinStack{Elements: []int{}, Min: []int{}}
}


func (this *MinStack) Push(x int)  {
    this.Elements = append(this.Elements, x)
    if len(this.Min) == 0 {
        this.Min = append(this.Min, x)
    } else if  x <= this.Min[len(this.Min)-1] {
        this.Min = append(this.Min, x)
    }
}


func (this *MinStack) Pop()  {    
	if this.Elements[len(this.Elements)-1] == this.Min[len(this.Min)-1] {
		this.Min = this.Min[:len(this.Min)-1]
	}
    this.Elements = this.Elements[:len(this.Elements)-1]
}


func (this *MinStack) Top() int {
    return this.Elements[len(this.Elements)-1]
}


func (this *MinStack) GetMin() int {
    return this.Min[len(this.Min)-1]
}