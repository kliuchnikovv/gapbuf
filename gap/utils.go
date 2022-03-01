package gap

func (g Gap) sizeOfData() int {
	return len(g.Data) - g.size
}

func (g Gap) LastIndex() int {
	return g.offset + g.size
}

func (g *Gap) moveGap(cursor int) {
	if cursor < g.offset {
		var bytes = make([]byte, g.offset-cursor)
		copy(bytes, g.Data[cursor:g.offset])

		for i, j := len(bytes)-1, 0; i >= 0; i, j = i-1, j+1 {
			g.Data[g.LastIndex()-1-j] = bytes[i]
		}
	} else {
		var bytes = make([]byte, cursor+g.size-g.LastIndex())
		copy(bytes, g.Data[g.LastIndex() : cursor+g.size])

		for i := 0; i < len(bytes); i++ {
			g.Data[g.offset+i] = bytes[i]
		}
	}

	g.offset = cursor

	// Unnecessary
	for i := 0; i < g.size; i++ {
		g.Data[g.offset+i] = 0
	}
}
