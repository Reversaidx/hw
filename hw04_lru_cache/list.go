package hw04lrucache

type (
	List interface {
		Len() int
		Front() *ListItem
		Back() *ListItem
		PushFront(v interface{}) *ListItem
		PushBack(v interface{}) *ListItem
		Remove(i *ListItem)
		MoveToFront(i *ListItem)
	}
)

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func (l list) Len() int {
	return l.len
}

func (l list) Back() *ListItem {
	return l.back
}

func (l list) Front() *ListItem {
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	tmp := &ListItem{Value: v, Next: nil, Prev: l.back}
	if l.len >= 1 {
		l.back.Next = tmp
		l.back = tmp
	} else {
		l.back = tmp
		l.front = tmp
	}

	l.len++
	return tmp
}

func (l *list) PushFront(v interface{}) *ListItem {
	tmp := &ListItem{Value: v, Next: l.front, Prev: nil}
	if l.len >= 1 {
		l.front.Prev = tmp
		l.front = tmp
	} else {
		l.back = tmp
		l.front = tmp
	}
	l.len++
	return tmp
}

func (l *list) Remove(i *ListItem) {

	if i.Prev != nil && i.Next != nil {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	} else if i.Prev == nil && i.Next != nil {
		i.Next.Prev = nil
		l.front = i.Next
	} else if i.Prev != nil && i.Next == nil {
		i.Prev.Next = nil
		l.back = i.Prev
	} else if i.Prev == nil && i.Next == nil {
		l.back = nil
		l.front = nil
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	tmp := i.Value
	l.Remove(i)
	l.PushFront(tmp)
}

type list struct {
	List  // Remove me after realization.
	front *ListItem
	back  *ListItem
	len   int
}

func NewList() List {
	return new(list)
}
