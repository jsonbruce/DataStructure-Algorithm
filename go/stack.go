package main

import (
	"fmt"
	"log"
)


type LinearStack struct {
	Data []int
}

func NewLinearStack() *LinearStack {
	ls := &LinearStack{make([]int, 0, 32)}
	return ls
}

func (this *LinearStack) Push(data int) {
	log.Println("Push:", data)
	this.Data = append(this.Data, data)
}

func (this *LinearStack) Pop() interface{} {
	if len(this.Data) != 0{
		d := this.Data[len(this.Data) - 1]
		this.Data = this.Data[:len(this.Data) - 1]
		log.Println("Pop:", d)
		return d
	}
	return nil
}

func (this *LinearStack) Print() {
	if len(this.Data) == 0 {
		fmt.Println("Empty Stack.")
	} else {
		for i := len(this.Data) - 1; i >= 0; i-- {
			fmt.Print("->", this.Data[i])
		}
		fmt.Println()
	}
}


type Node struct {
	Data interface{}
	Next *Node
}

type LinkedStack struct {
	Top *Node
}

func NewLinkedStack() *LinkedStack {
	ls := &LinkedStack{nil}
	return ls
}

func (this *LinkedStack) Push(data interface{}) {
	log.Println("Push", data)
	newTop := &Node{data, nil}
	newTop.Next = this.Top
	this.Top = newTop
}

func (this *LinkedStack) Pop() interface{} {
	if this.Top != nil {
		data := this.Top.Data
		this.Top = this.Top.Next
	    log.Println("Pop", data)
		return data
	}
	return nil
}

func (this *LinkedStack) Print() {
	if this.Top == nil {
		fmt.Println("Empty Stack.")
	} else {
		top := this.Top
		s := ""
		for top != nil {
			s += "->" + fmt.Sprintf("%d", top.Data)
			top = top.Next
		}
		fmt.Println(s)
	}
}


func main() {
	fmt.Println("LinearStack")
	s := NewLinearStack()
	s.Push(12)
	s.Push(14)
	s.Push(16)
	s.Print()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Print()


	fmt.Println("\nLinkedStack")
	ls := NewLinkedStack()
	ls.Push(12)
	ls.Push(14)
	ls.Push(16)
	ls.Print()
	ls.Pop()
	ls.Pop()
	ls.Pop()
	ls.Print()
}
