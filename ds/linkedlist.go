package ds

type Node[T any] struct {
	value      int
	next, prev *Node[T]
}

type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
} 
