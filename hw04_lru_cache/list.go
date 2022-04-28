package hw04lrucache

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
	Prev  *ListItem
	Next  *ListItem
}

type list struct {
	firstNode *ListItem
	lastNode  *ListItem
	len       int
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem { // Первый эл
	if l.len == 0 {
		return nil
	}
	return l.firstNode
}

func (l list) Back() *ListItem { // последний эл
	if l.len == 0 {
		return nil
	}
	return l.lastNode
}

func (l *list) PushBack(v interface{}) *ListItem { // в конец
	var newEl ListItem
	newEl.Value = v

	if l.Front() == nil { // Пустой список
		l.firstNode = &newEl
		l.lastNode = &newEl
		l.len = 1
	} else {
		newEl.Prev = l.lastNode
		l.lastNode.Next = &newEl
		l.lastNode = &newEl
		l.len++
	}
	return &newEl
}

func (l *list) PushFront(v interface{}) *ListItem { // в начало
	var newEl ListItem
	newEl.Value = v

	if l.Front() == nil { // Пустой список
		l.firstNode = &newEl
		l.lastNode = &newEl
		l.len = 1
	} else {
		newEl.Next = l.firstNode
		l.firstNode.Prev = &newEl
		l.firstNode = &newEl
		l.len++
	}
	return &newEl
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil { // Первый
		return
	}

	if i.Next == nil { // последний
		l.lastNode = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	i.Prev.Next = i.Next
	i.Next = l.firstNode
	l.firstNode.Prev = i
	l.firstNode = i
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil { // Первый
		l.firstNode = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil { // последний
		l.lastNode = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.len--
}

func NewList() List {
	return new(list)
}
