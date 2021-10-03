package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	size  int
	front *ListItem
	back  *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) PushFront(v interface{}) *ListItem {
	listItem := ListItem{
		Value: v,
	}

	if l.size == 0 {
		l.back = &listItem
	} else {
		fmt.Println(l.front)
		l.front.Prev = &listItem
		listItem.Next = l.front
	}

	l.front = &listItem
	l.size++

	return &listItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	listItem := ListItem{
		Value: v,
	}

	if l.size == 0 {
		l.front = &listItem
	} else {
		l.back.Next = &listItem
		listItem.Prev = l.back
	}

	l.back = &listItem
	l.size++

	return &listItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		if i.Next != nil {
			l.front = i.Next
		} else {
			l.front = nil
		}
	}
	if i.Next == nil {
		if i.Prev != nil {
			l.back = i.Prev
		} else {
			l.back = nil
		}
	}

	if i.Prev != nil {
		if i.Next != nil {
			i.Prev.Next = i.Next
			i.Next.Prev = i.Prev
		} else {
			i.Prev.Next = nil
			l.back = i.Prev
		}
	} else if i.Next != nil {
		i.Next.Prev = nil
		l.front = i.Next
	}

	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if *i == *l.front {
		return
	}

	if i.Next == nil && i.Prev != nil {
		l.back = i.Prev
		l.back.Next = nil
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if l.front != nil {
		i.Next = l.front
	}

	i.Prev = nil
	l.front.Prev = i
	l.front = i
}
