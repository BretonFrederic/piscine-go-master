package piscine

func ListSize(l *List) int {
	count := 0
	for n := l.Head; n != nil; n = n.Next {
		count++
	}
	return count
}
