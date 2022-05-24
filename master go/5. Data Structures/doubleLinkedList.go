package main

import ("fmt"
)

type Node struct {
	Value int
	Previous *Node
	Next *Node
}

func addNode(t * Node, v int) int {
	if root = nil {
		t = &Node{v, nil, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		 return -1
	}

	if t.Next == nil {
		temp := t

		t.Next = &Node{v, temp, nil}
		return -2
	}
	return addNode(t.Next, v)
}

func traverse (t * Node){
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}

	temp := t 
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("%d -> ", temp.Valie)
	fmt.Println
}