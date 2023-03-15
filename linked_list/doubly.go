package linked_list

type dNode struct {
	prev *dNode
	next *dNode
	key  any
}

type DoublyLL struct {
	head *dNode
	tail *dNode
}

func (d *DoublyLL) returnList() []any {
	list := make([]any, 0)
	current := d.head
	for current != nil {
		list = append(list, current.key)
		current = current.next
	}
	return list
}

func MakeDoublyLL() *DoublyLL {
	return &DoublyLL{nil, nil}
}

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
