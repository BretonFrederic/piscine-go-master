package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	current := l.Head
	for n := l.Head; n != nil; n = n.Next {
		current = n
	}
	return current.Data
}
