package linkedlist

import "fmt"

// This modifies the head pointer in the calling function
func InsertAtEndInPlace(head **Node, x int) {
	fmt.Println("\nIn InsertAtEndInPlace:")
	tmp := Node{Next: nil, Data: x}
	fmt.Println("Inserting", tmp, "at the end")
	if *head == nil {
		*head = &tmp
	} else {
		var p *Node
		p = *head
		for ; p.Next != nil; p = p.Next {
		}
		p.Next = &tmp
	}
	fmt.Printf("Address of head node %p\n", &**head)
	fmt.Println("Address of head pointer ", &*head)
	fmt.Println("Returning")
}

// This returns the head
func InsertAtEnd(head *Node, x int) *Node {
	fmt.Println("\nIn InsertAtEnd:")
	tmp := Node{Next: nil, Data: x}
	fmt.Println("Inserting", tmp, "at the end")
	current := head
	if head == nil {
		head = &tmp
	} else {
		for ; current.Next != nil; current = current.Next {
		}
		current.Next = &tmp
	}
	fmt.Printf("Address of node %p\n", &*head)
	fmt.Println("Address of head pointer ", &head)
	fmt.Println("Returning")
	return head
}
func InsertAtStart(head **Node, x int) {
	tmp := Node{Data: x, Next: nil}
	fmt.Println("\nInserting", tmp, "at the start")
	current := *head
	if current == nil {
		current = &tmp
		return
	}
	tmp.Next = current
	*head = &tmp
}

func InsertAtP(head **Node, p int, x int) {
	current := *head
	if current == nil {
		fmt.Println("List is empty")
		return
	}
	for i := 1; i < p; i++ {
		if current == nil {
			fmt.Println("Position Not Found")
			return
		}
		current = current.Next
	}
	tmp := Node{Data: x, Next: nil}
	fmt.Println("\nInserting", tmp, "at the position", p)
	tmp.Next = current.Next
	current.Next = &tmp
}
