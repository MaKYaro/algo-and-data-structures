package linked_list

import "fmt"

// node is an element of SinglyLL
type node struct {
	next *node
	key  any
}

// noSuchElement occurs when there is no elem in the linked list
type noSuchElement struct {
	elem any
}

func (n noSuchElement) Error() string {
	return fmt.Sprintf("no such element: %v", n.elem)
}

// SinglyLL is a singly linked list
type SinglyLL struct {
	head *node
	tail *node
}

// returnList returns array of all elements keys of SinglyLL
// in order they are in the SinglyLL starting with the head of the list
func (s *SinglyLL) returnList() []any {
	current := s.head
	list := make([]any, 0)
	for current != nil {
		list = append(list, current.key)
		current = current.next
	}
	return list
}

// MakeSinglyLL returns new entity of SinglyLL
func MakeSinglyLL() *SinglyLL {
	return &SinglyLL{nil, nil}
}

// InsertHead inserts new element with the key in the head place of SinglyLL
func (s *SinglyLL) InsertHead(key any) {
	switch s.head {
	case nil:
		n := &node{nil, key}
		s.head = n
		s.tail = n
	default:
		nextAfterHead := s.head
		s.head = &node{nextAfterHead, key}
	}
}

// InsertTail inserts new element with the key in the tail place of SinglyLL
func (s *SinglyLL) InsertTail(key any) {
	switch s.head {
	case nil:
		n := &node{nil, key}
		s.head = n
		s.tail = n
	default:
		tail := &node{nil, key}
		beforeTail := s.tail
		beforeTail.next = tail
		s.tail = tail
	}
}

// Delete deletes first element with the key from SinglyLL
// error is nil if there is element with the key or noSuchError otherwise
func (s *SinglyLL) Delete(key any) error {
	switch {
	case s.head == nil:
		return noSuchElement{key}
	case s.head.key == key:
		s.head = s.head.next
		return nil
	default:
		previous := s.head
		current := s.head.next
		for current != nil && current.key != key {
			previous = current
			current = current.next
		}
		//if current == nil {
		//	return noSuchElement{key}
		//}
		//previous.next = current.next
		//return nil
		switch {
		case current == nil:
			return noSuchElement{key}
		case current.next == nil:
			s.tail = previous
			previous.next = nil
			return nil
		default:
			previous.next = current.next
			return nil
		}
	}
}

// Search returns sequence number of the first element with the key and nil error
// -1 and noSuchElement error if there is no such element in the list
func (s *SinglyLL) Search(key any) (int, error) {
	idx := 0
	current := s.head
	for current != nil && current.key != key {
		current = current.next
		idx++
	}
	if current == nil {
		return -1, noSuchElement{key}
	}
	return idx, nil
}
