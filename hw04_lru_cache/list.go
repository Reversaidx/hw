package hw04lrucache

import (
	"errors"
)

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

var ErrInvalidSrtDst = errors.New("Dst index has to be more than srt")
var ListMap = make(map[*ListItem]int)

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func (l list) Len() int {
	return len(l.ListItems)
}
func (l list) Back() *ListItem {
	if l.Len() > 0 {
		return l.ListItems[len(l.ListItems)-1]
	}
	return nil
}
func (l list) Front() *ListItem {
	if l.Len() > 0 {
		return l.ListItems[0]
	}
	return nil
}
func (l *list) PushBack(v interface{}) *ListItem {
	tmp := &ListItem{Value: v, Next: nil, Prev: nil}
	if l.Len() > 0 {
		l.ListItems = append(l.ListItems, tmp)
		l.bind(l.Len()-2, l.Len()-1)
		l.UpdateMap(tmp, l.Len()-1-l.offset)
	} else {
		l.ListItems = append(l.ListItems, tmp)
		l.UpdateMap(tmp, 0)
	}
	return tmp
}

func (l *list) PushFront(v interface{}) *ListItem {
	tmp := &ListItem{Value: v, Next: nil, Prev: nil}
	if l.Len() > 0 {
		tmpArr := make([]*ListItem, l.Len()+1)
		tmpArr = append(tmpArr, tmp)

		tmpArr = append(tmpArr, l.ListItems[:])
		tmpArr[0] = tmp
		l.bind(0, 1)

	} else {
		l.ListItems = append(l.ListItems, tmp)
	}

	l.UpdateMap(tmp, 0)
	l.offset++
	return tmp
}

func (l *list) Remove(i *ListItem) {
	pos := l.Maps[i] + l.offset
	l.bind(pos-1, pos+1)
	l.ListItems = append(l.ListItems[:pos], l.ListItems[pos+1:]...)

}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i)

}

func (l *list) bind(srtPos, dstPos int) {
	if srtPos > dstPos {
		srtPos, dstPos = dstPos, srtPos
	}
	l.ListItems[srtPos].Next = l.ListItems[dstPos]
	l.ListItems[dstPos].Prev = l.ListItems[srtPos]
}

type list struct {
	List      // Remove me after realization.
	ListItems []*ListItem
	Maps      map[*ListItem]int
	offset    int

	// Place your code here.
}

func (l *list) UpdateMap(i *ListItem, pos int) {
	if l.Maps == nil {
		l.Maps = make(map[*ListItem]int)
	}
	l.Maps[i] = pos
}
func NewList() List {
	return new(list)
}
