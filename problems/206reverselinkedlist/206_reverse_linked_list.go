package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	values := []int{}

	current := head

	for {
		if current == nil {
			break
		}

		values = append(values, current.Val)
		current = current.Next
	}

	return listToReversedLinkedList(values)
}

func listToReversedLinkedList(l []int) *ListNode {
	var root *ListNode
	var last *ListNode

	for i := range l {
		n := l[len(l)-i-1]

		node := &ListNode{
			Val: n,
		}

		if root == nil {
			root = node
		} else {
			last.Next = node
		}

		last = node
	}

	return root
}

func listToLinkedList(l []int) *ListNode {
	var root *ListNode
	var last *ListNode

	for _, n := range l {
		node := &ListNode{
			Val: n,
		}

		if root == nil {
			root = node
		} else {
			last.Next = node
		}

		last = node
	}

	return root
}
