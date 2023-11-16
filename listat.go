package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	i := l
	inc := 0
	for i != nil {
		if pos == inc {
			return i
		}
		inc++
		i = i.Next
	}
	return nil
}
