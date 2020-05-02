package main

func main() {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}
	l3 := &ListNode{Val: 1}
	listNodes := []*ListNode{l, l1, l2, l3}

	lists := mergeKLists(listNodes)
	println(lists.Val)
	println(lists.Next.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	if a == nil || b == nil {
		if a == nil {
			return b
		}
		return a
	}
	var head ListNode
	tail := &head
	aPtr := a
	bPtr := b
	for aPtr != nil && bPtr != nil {
		if aPtr.Val < bPtr.Val {
			tail.Next = aPtr
			aPtr = aPtr.Next
		} else {
			tail.Next = bPtr
			bPtr = bPtr.Next
		}
		tail = tail.Next
	}
	if aPtr == nil {
		tail.Next = bPtr
	} else {
		tail.Next = aPtr
	}
	return head.Next
}
