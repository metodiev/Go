package main

import (
	"bufio"
	"fmt"
)

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func (doublyLinkedList *DoublyLinkedList) insertNodeIntoDoublyLinkedList(nodeData int32) {
	node := &DoublyLinkedListNode{
		next: nil,
		prev: nil,
		data: nodeData,
	}

	if doublyLinkedList.head == nil {
		doublyLinkedList.head = node
	} else {
		doublyLinkedList.tail.next = node
		node.prev = doublyLinkedList.tail
	}

	doublyLinkedList.tail = node
}

func printDoublyLinkedList(node *DoublyLinkedListNode, sep string, writer *bufio.Writer) {
	for node != nil {
		fmt.Fprintf(writer, "%d", node.data)

		node = node.next

		if node != nil {
			fmt.Fprintf(writer, sep)
		}
	}
}

/*
 * Complete the 'sortedInsert' function below.
 *
 * The function is expected to return an INTEGER_DOUBLY_LINKED_LIST.
 * The function accepts following parameters:
 *  1. INTEGER_DOUBLY_LINKED_LIST llist
 *  2. INTEGER data
 */

/*
 * For your reference:
 *
 * DoublyLinkedListNode {
 *     data int32
 *     next *DoublyLinkedListNode
 *     prev *DoublyLinkedListNode
 * }
 *
 */

func sortedInsert(llist *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	newNode := DoublyLinkedListNode{}
	newNode.data = data

	// if llist == nil {
	// 	return &newNode
	// }

	currentNode := llist
	previous := DoublyLinkedListNode{}

	if data < currentNode.data {
		newNode.next = currentNode
		newNode.prev = nil
		currentNode.prev = &newNode
		return &newNode
	}

	for currentNode != nil && data > currentNode.data {
		previous.data = currentNode.data
		previous.next = currentNode.next
		currentNode = currentNode.next
	}

	if currentNode == nil {
		newNode.prev = &previous
		newNode.next = nil
		previous.next = &newNode
	} else {
		newNode.prev = &previous
		newNode.next = previous.next
		previous.next = &newNode
		newNode.next.prev = &newNode
	}

	return llist
}

func main() {

	llist := DoublyLinkedList{}
	llist.insertNodeIntoDoublyLinkedList(1)
	llist.insertNodeIntoDoublyLinkedList(2)
	llist.insertNodeIntoDoublyLinkedList(3)

	sortedInsert(llist.head, 4)

	//printDoublyLinkedList(llist1, " ", writer)

}
