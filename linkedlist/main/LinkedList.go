package main

import (
	"fmt"
	"sync"
)

type DoubleObject interface {
}

type DoubleNode struct {
	Data DoubleObject
	Prev *DoubleNode
	Next *DoubleNode
}

type DoubleList struct {
	mutex *sync.RWMutex
	Size  uint
	Head  *DoubleNode
	Tail  *DoubleNode
}

func (list *DoubleList) Init() {
	list.mutex = new(sync.RWMutex)
	list.Size = 0
	list.Head = nil
	list.Tail = nil
}

/**
根据下标获取
*/
func (list *DoubleList) Get(index uint) *DoubleNode {
	if list.Size == 0 || index > list.Size-1 {
		return nil
	}
	if index == 0 {
		return list.Head
	}
	node := list.Head
	var i uint
	for i = 1; i <= index; i++ {
		node = node.Next
	}
	return node
}

/**
追加
*/
func (list *DoubleList) Append(node *DoubleNode) bool {
	if node == nil {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		node.Next = nil
		node.Prev = nil
	} else {
		node.Prev = list.Tail
		node.Next = nil
		list.Tail.Next = node
		list.Tail = node
	}
	list.Size++
	return true
}

/**
插入
*/
func (list *DoubleList) Insert(index uint, node *DoubleNode) bool {
	if index > list.Size || node == nil {
		return false
	}
	if index == list.Size {
		return list.Append(node)
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		node.Next = list.Head
		list.Head = node
		list.Head.Prev = nil
		list.Size++
		return true
	}

	nextNode := list.Get(index)
	node.Prev = nextNode.Prev
	node.Next = nextNode
	nextNode.Prev.Next = node
	nextNode.Prev = node
	list.Size++
	return true
}

/**
删除
*/
func (list *DoubleList) Delete(index uint) bool {
	if index > list.Size-1 {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		if list.Size == 1 {
			list.Head = nil
			list.Tail = nil
		} else {
			list.Head.Next.Prev = nil
			list.Head = list.Head.Next
		}
		list.Size--
		return true
	}
	if index == list.Size-1 {
		list.Tail.Prev.Next = nil
		list.Tail = list.Tail.Prev
		list.Size--
		return true
	}
	node := list.Get(index)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Size--
	return true
}

/**
打印输出
*/
func (list *DoubleList) Display() {
	if list == nil || list.Size == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this double list size is %d \n", list.Size)
	ptr := list.Head
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Next
	}
}

/**
颠倒打印
*/
func (list *DoubleList) Reverse() {
	if list == nil || list.Size == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this double list size is %d \n", list.Size)
	ptr := list.Tail
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Prev
	}
}

/**
双向链表翻转
*/
func (list *DoubleList) flip() {
	if list == nil || list.Size == 0 {
		fmt.Println("this double list is nil")
		return
	}

	if list.Size == 1 {
		fmt.Printf("this node is %d", list.Head)
	}
	node := list.Head
	list.Head, list.Tail = list.Tail, list.Head

	var i uint
	for i = 0; i < list.Size; i++ {
		node.Prev, node.Next = node.Next, node.Prev
		node = node.Prev
	}

}

/**
 * 单向链表
 */
type Node struct {
	data int
	next *Node
}

func ShowNode(p *Node) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func main() {
	var head = new(Node)
	head.data = 1
	var node1 = new(Node)
	node1.data = 2

	head.next = node1
	var node2 = new(Node)
	node2.data = 3

	node1.next = node2
	ShowNode(head)

	var list = new(DoubleList)
	list.Init()
	for i := 0; i < 10; i++ {
		node := new(DoubleNode)
		node.Data = i
		list.Append(node)
	}
	list.Display()
	list.flip()
	list.Display()

}
