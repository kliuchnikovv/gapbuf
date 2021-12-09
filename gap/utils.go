package gap

func (g Gap) sizeOfData() int {
	return len(*g.data) - g.size
}

func (g Gap) LastIndex() int {
	return g.offset + g.size
}

func (g *Gap) moveGap() {
	if g.cursor < g.offset {
		var bytes = (*g.data)[g.cursor:g.offset]

		for i, char := range bytes {
			(*g.data)[g.cursor+g.size+i], (*g.data)[g.cursor+i] = char, (*g.data)[g.cursor+g.size+i]
		}
	} else {
		var bytes = (*g.data)[g.offset+g.size : g.cursor+g.size]

		for i, char := range bytes {
			(*g.data)[g.offset+i], (*g.data)[g.offset+g.size+i] = char, (*g.data)[g.offset+i]
		}
	}

	g.offset = g.cursor
}

///////////////////////// a             c
//func (g *Gap) moveBytes(offset, size, secondOffset int) {
//	var bytes = (*g.data)[offset:secondOffset]
//
//	for i, char := range bytes {
//		(*g.data)[offset+size+i], (*g.data)[offset+i] = char, (*g.data)[offset+size+i]
//	}
//}
