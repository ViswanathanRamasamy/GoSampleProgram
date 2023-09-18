/*
LinkList: sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
Author Name: Viswanathan Ramasamy
Email id: rviswawt@gmail.com
*/
====================================================
Linklist:
for loop can also be replaced like:

        curr := ll.head
        for ;curr.next != nil ;curr = curr.next {
            
        }
=================
 //struct initialization { data: data , next : nil } not =
 // there no -> for the pointer member check, it should be .
 //current.next can be replaced by (*current).next since current is the pointer
==================		
package main

import (
	"fmt"
)

// Node represents a node in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	head *Node
}

// Append adds a new node with the given data to the end of the list
func (ll *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next  //or (*current).next
	}
	current.next = newNode
}

// Display prints the elements of the linked list
func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := LinkedList{}

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)

	fmt.Println("Linked List elements:")
	ll.Display()
}
==================================================================
reverse the linklist in go program:
package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) Reverse() {
	var prev, next *Node
	current := ll.head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	ll.head = prev
}

func main() {
	ll := LinkedList{}

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)

	fmt.Println("Original Linked List:")
	ll.Display()

	ll.Reverse()

	fmt.Println("Reversed Linked List:")
	ll.Display()
}
========================================