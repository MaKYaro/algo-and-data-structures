package linked_list

// rNode is an element of RingLL
// isTail determines if the element is the tail
type rNode struct {
	isTail bool
	next   *rNode
	key    any
}

// RingLL is ring linked list
// tail.next points to head element
type RingLL struct {
	head *rNode
	tail *rNode
}

// MakeRingLL returns new entity of RingLL
func MakeRingLL() *RingLL {
	return &RingLL{nil, nil}
}

func (r *RingLL) returnList() []any {
	current := r.head
	list := make([]any, 0)
	for current != nil && !current.isTail {
		list = append(list, current.key)
		current = current.next
	}
	if current != nil {
		list = append(list, current.key)
	}
	return list
}

// InsertHead method inserts new element with the key in the head place of RingLL
func (r *RingLL) InsertHead(key any) {
	switch r.head {
	case nil:
		n := &rNode{true, nil, key}
		n.next = n
		r.head = n
		r.tail = n
	default:
		n := &rNode{false, r.head, key}
		r.head = n
		r.tail.next = n
	}
}

// InsertTail method inserts new element with the key in the tail place of RingLL
func (r *RingLL) InsertTail(key any) {
	switch r.head {
	case nil:
		n := &rNode{true, nil, key}
		n.next = n
		r.head = n
		r.tail = n
	default:
		n := &rNode{true, r.head, key}
		r.tail.next = n
		r.tail.isTail = false
		r.tail = n
	}
}

// Delete deletes first element with the key from RingLL
// error is nil if there is element with the key or noSuchError otherwise
func (r *RingLL) Delete(key any) error {
	switch {
	case r.head == nil:
		return noSuchElement{key}
	case r.head.key == key:
		r.head = r.head.next
		r.tail.next = r.head
		return nil
	default:
		previous := r.head
		current := r.head.next
		for current.isTail != true && current.key != key {
			previous = current
			current = current.next
		}
		switch {
		case current.isTail && current.key != key:
			return noSuchElement{key}
		case current.isTail:
			previous.next = r.head
			previous.isTail = true
			r.tail = previous
			return nil
		default:
			previous.next = current.next
			return nil
		}
	}
}

// Search returns sequence number of the first element with the key and nil error
// -1 and noSuchElement error if there is no such element in the list
func (r *RingLL) Search(key any) (int, error) {
	idx := 0
	current := r.head
	if current == nil {
		return -1, noSuchElement{key}
	}
	for current.isTail != true && current.key != key {
		current = current.next
		idx++
	}
	switch {
	case current.isTail && current.key != key:
		return -1, noSuchElement{key}
	case current.isTail:
		return idx, nil
	}
	return idx, nil
}
