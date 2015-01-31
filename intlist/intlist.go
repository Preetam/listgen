// Generated by listgen. DO NOT EDIT.
// Command: listgen -package intlist -type int -cmp "return a-b"

package intlist

import (
	"errors"
	"fmt"
	"sync/atomic"
	"unsafe"
)

var (
	ErrIntListValueExists   = errors.New("intlist (IntList): value exists")
	ErrIntListValueNotFound = errors.New("intlist (IntList): value does not exist")
)

type IntList struct {
	head unsafe.Pointer
}

type IntListNode struct {
	val  int
	next unsafe.Pointer
}

type IntListIterator struct {
	list    *IntList
	current *IntListNode
	valid   bool
}

func IntListCmp(a, b int) int {
	return a - b
}

// NewList returns a lock-free ordered list with values of type int.
func NewList() *IntList {
	return &IntList{}
}

// Insert inserts v into the list in order. An error is returned if v is already present.
func (l *IntList) Insert(v int) error {
	n := &IntListNode{
		val:  v,
		next: nil,
	}

HEAD:
	headPtr := atomic.LoadPointer(&l.head)

	if headPtr == nil {
		if !atomic.CompareAndSwapPointer(&l.head, headPtr, unsafe.Pointer(n)) {
			goto HEAD
		}

		return nil
	}

	headNode := (*IntListNode)(headPtr)
	if IntListCmp(headNode.val, n.val) > 0 {
		n.next = headPtr
		if !atomic.CompareAndSwapPointer(&l.head, headPtr, unsafe.Pointer(n)) {
			goto HEAD
		}

		return nil
	}

NEXT:
	nextPtr := atomic.LoadPointer(&headNode.next)
	if nextPtr == nil {
		if !atomic.CompareAndSwapPointer(&headNode.next, nextPtr, unsafe.Pointer(n)) {
			goto NEXT
		}

		return nil
	}

	nextNode := (*IntListNode)(nextPtr)
	if IntListCmp(nextNode.val, n.val) > 0 {
		n.next = nextPtr
		if !atomic.CompareAndSwapPointer(&headNode.next, nextPtr, unsafe.Pointer(n)) {
			goto NEXT
		}

		return nil
	}

	if IntListCmp(nextNode.val, n.val) == 0 {
		return ErrIntListValueExists
	}

	headNode = nextNode
	goto NEXT
}

// Remove removes v from the list. An error is returned if v is not present.
func (l *IntList) Remove(v int) error {
HEAD:
	headPtr := atomic.LoadPointer(&l.head)

	if headPtr == nil {
		return ErrIntListValueNotFound
	}

	headNode := (*IntListNode)(headPtr)

	if IntListCmp(headNode.val, v) == 0 {
		nextPtr := atomic.LoadPointer(&headNode.next)
		if !atomic.CompareAndSwapPointer(&l.head, headPtr, nextPtr) {
			goto HEAD
		}

		return nil
	}

NEXT:
	nextPtr := atomic.LoadPointer(&headNode.next)
	if nextPtr == nil {
		return ErrIntListValueNotFound
	}

	nextNode := (*IntListNode)(nextPtr)

	if IntListCmp(nextNode.val, v) > 0 {
		return ErrIntListValueNotFound
	}

	if IntListCmp(nextNode.val, v) == 0 {
		replacementPtr := atomic.LoadPointer(&nextNode.next)
		if !atomic.CompareAndSwapPointer(&headNode.next, nextPtr, replacementPtr) {
			goto NEXT
		}

		return nil
	}

	headNode = nextNode
	goto NEXT
}

// NewIterator returns a new iterator. Values can be read
// after Next is called.
func (l *IntList) NewIterator() *IntListIterator {
	return &IntListIterator{
		list:  l,
		valid: true,
	}
}

// Next positions the iterator at the next node in the list.
// Next will be positioned at the head on the first call.
// The return value will be true if a value can be read from the list.
func (i *IntListIterator) Next() bool {
	if !i.valid {
		return false
	}

	if i.current == nil {
		if i.list.head == nil {
			i.valid = false
			return false
		}

		i.current = (*IntListNode)(i.list.head)
		return true
	}

	i.current = (*IntListNode)(i.current.next)

	i.valid = i.current != nil
	return i.valid
}

// Value reads the value from the current node of the iterator.
// An error is returned if a value cannot be retrieved.
func (i *IntListIterator) Value() (int, error) {
	var v int

	if i.current == nil {
		return v, ErrIntListValueNotFound
	}

	return i.current.val, nil
}

// String returns the string representation of the list.
func (l *IntList) String() string {
	output := ""

	if l.head == nil {
		return output
	}

	i := l.NewIterator()

	for i.Next() {
		v, _ := i.Value()
		output += fmt.Sprintf("%v ", v)
	}

	return output
}
