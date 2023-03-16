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
