package main


/*
{
	from:"https://leetcode-cn.com/problems/add-two-numbers/",
	reference:[],
	description:"两数相加",
	solved:true,
}
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res = l2
	var rest = (l1.Val + l2.Val) / 10
	l2.Val = (l1.Val + l2.Val) % 10
	if l1.Next == nil && l2.Next == nil && rest != 0 {
		l2.Next = new(ListNode)
		l2.Next.Val = rest
		l2.Next.Next = nil
		return res
	}

	for !(l1.Next == nil && l2.Next == nil) {
		if l1.Next != nil && l2.Next != nil {
			sum := l1.Next.Val + l2.Next.Val + rest
			l2.Next.Val = sum % 10
			rest = sum / 10
			l1 = l1.Next
			l2 = l2.Next
		} else if l1.Next == nil {
			sum := l2.Next.Val + rest
			l2.Next.Val = sum % 10
			rest = sum / 10
			//go on
			l2 = l2.Next
		} else { //l2.Next==nil
			sum := l1.Next.Val + rest
			rest = sum / 10
			l2.Next = new(ListNode)
			l2.Next.Val = sum % 10
			l2.Next.Next = nil
			//go on
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	if rest != 0 {
		l2.Next = new(ListNode)
		l2.Next.Val = rest
		l2.Next.Next = nil
	}
	return res
}

func main() {

}
