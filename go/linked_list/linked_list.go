package linked_list

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
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
	for cur != nil {
		tmp := cur.next
		if tmp == nil {
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

func (this *LinkedList) ReverseShort() {
	if this.head == nil || this.head.next == nil {
		return
	}

	var next *ListNode
	cur := this.head
	for cur.next != nil {
		next = cur.next
		cur.next = next.next
		next.next = this.head
		this.head = next
	}

	// if tail exists => this.tail = cur
}

func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	cur := this.head
	var pre *ListNode
	for cur != nil && cur != p {
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	if pre == nil {
		this.head = p.next
	} else {
		pre.next = p.next
	}
	p = nil
	this.length--
	return true
}

func (this *LinkedList) DeleteNode2(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := &this.head
	// if p is in linked list, '*cur != nil' can be removed
	for *cur != nil && *cur != p {
		cur = &(*cur).next
	}
	if nil == *cur {
		return false
	}
	*cur = p.next
	this.length--
	return true
}
