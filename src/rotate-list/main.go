package main

/*
	{
		from:"https://leetcode-cn.com/problems/rotate-list",
		reference:[],
		description:"旋转list",
		solved:true,
	}
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	length := 1
	for p.Next != nil {
		length++
		p = p.Next
	}
	p.Next = head
	k %= length
	k = length - k
	for k > 0 {
		p = p.Next
		k--
	}
	res := p.Next
	p.Next = nil
	return res
}
