package _7_linkedlist

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head *ListNode
}

func (this LinkedList) Print() {
	cur := this.head
	format := ""

	for cur != nil {
		format += fmt.Sprintf("%+v", cur.value)
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

func (this *LinkedList) Reverse() {
	if this.head == nil || this.head.next == nil {
		return
	}

	var pre *ListNode
	cur := this.head
	for cur != nil  {
		tmp := cur.next
		if tmp ==nil{
			cur.next = pre
			pre = cur
			break
		}
		cur.next = pre
		pre = cur
		cur = tmp
	}

	this.head = pre
}



