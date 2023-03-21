package linked_list

// dNode is an element of DoublyLL
type dNode struct {
	prev *dNode
	next *dNode
	key  any
}

// DoublyLL is doubly linked list
type DoublyLL struct {
	head *dNode
	tail *dNode
}

// returnList returns array of all elements keys of DoublyLL
// in order they are in the DoublyLL starting with the head of the list
func (d *DoublyLL) returnList() []any {
	list := make([]any, 0)
	current := d.head
	for current != nil {
		list = append(list, current.key)
		current = current.next
	}
	return list
}

// MakeDoublyLL returns new entity of DoublyLL
func MakeDoublyLL() *DoublyLL {
	return &DoublyLL{nil, nil}
}

// InsertHead inserts new element with the key in the head place of DoublyLL
func (d *DoublyLL) InsertHead(key any) {
	switch d.head {
	case nil:
		n := &dNode{nil, nil, key}
		d.head = n
		d.tail = n
	default:
		nextAfterHead := d.head
		d.head = &dNode{nil, nextAfterHead, key}
		nextAfterHead.prev = d.head
	}
}

// InsertTail inserts new element with the key in the tail place of DoublyLL
func (d *DoublyLL) InsertTail(key any) {
	switch d.head {
	case nil:
		n := &dNode{nil, nil, key}
		d.head = n
		d.tail = n
	default:
		beforeTail := d.tail
		n := &dNode{beforeTail, nil, key}
		d.tail = n
		beforeTail.next = n
	}

}

// Delete deletes first element with the key from DoublyLL
// error is nil if there is element with the key or noSuchError otherwise
func (d *DoublyLL) Delete(key any) error {
	switch {
	case d.head == nil:
		return noSuchElement{key}
	case d.head.key == key:
		nextAfterHead := d.head
		d.head = nextAfterHead
		nextAfterHead.prev = nil
		return nil
	default:
		current := d.head.next
		for current != nil && current.key != key {
			current = current.next
		}
		switch {
		case current == nil:
			return noSuchElement{key}
		case current.next == nil:
			tail := current.prev
			tail.next = nil
			d.tail = tail
			return nil
		default:
			current.prev.next = current.next
			current.next.prev = current.prev
			return nil
		}
	}
}

// Search returns sequence number of the first element with the key and nil error
// -1 and noSuchElement error if there is no such element in the list
func (d *DoublyLL) Search(key any) (int, error) {
	idx := 0
	current := d.head
	for current != nil && current.key != key {
		current = current.next
		idx++
	}
	if current == nil {
		return -1, noSuchElement{key}
	}
	return idx, nil
}

// Reverse reverses DoublyLL
// head becomes tail, tail becomes head
// each dNode.prev has value of dNode.next otherwise dNode.next has value of dNode.prev
func (d *DoublyLL) Reverse() {
	switch d.head {
	case nil:
	default:
		tail := d.tail
		current := d.head
		for current.next != nil {
			next := current.next
			tmp := current.prev
			current.prev = current.next
			current.next = tmp
			current = next
		}
		tail.next = tail.prev
		tail.prev = nil
		d.head, d.tail = d.tail, d.head
	}

}
