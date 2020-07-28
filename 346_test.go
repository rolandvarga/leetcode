package main_test

import (
	"fmt"
	"testing"
)

type MovingAverage struct {
	Size  int
	Items []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{Size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	this.Items = append(this.Items, val)

	var sum int
	if len(this.Items) < this.Size {
		for i := 0; i < len(this.Items); i++ {
			sum = sum + this.Items[i]
		}
	} else if len(this.Items) == this.Size {
		window := len(this.Items) - this.Size
		for i := window; i < len(this.Items)-1; i++ {
			sum = sum + this.Items[i]
		}
	}
	return float64(sum / len(this.Items))
}

func TestMovingAverage(t *testing.T) {
	m := Constructor(3)
	fmt.Printf("got %v expected %s\n", m.Next(1), "1")
	fmt.Printf("got %v expected %s\n", m.Next(10), "(1 + 10) / 2")
	fmt.Printf("got %v expected %s\n", m.Next(3), "(1 + 10 + 3) / 3")
	fmt.Printf("got %v expected %s\n", m.Next(5), "(10 + 3 + 5) / 3")

}
