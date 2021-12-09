package gap

type Gap struct {
	offset int
	size   int

	data   *[]byte
	cursor int
}

func New(size int, pointer *[]byte) *Gap {
	return &Gap{
		offset: 0,
		size:   size,
		data:   pointer,
	}
}

func (g *Gap) Size() int {
	return g.size
}

func (g *Gap) SetCursor(cursor int) {
	if cursor < 0 || cursor > g.sizeOfData() {
		return
		//panic(fmt.Sprintf("cursor is out of range (len: %d, cursor: %d)", g.sizeOfData(), cursor))
	}

	if cursor > g.offset {
		g.cursor = cursor + g.size
	} else {
		g.cursor = cursor
	}

}

func (g Gap) GetCursor() int {
	if g.cursor > g.offset {
		return g.cursor - g.size
	}
	return g.cursor
}

func (g Gap) Offset() int {
	return g.offset
}

func (g *Gap) Insert(char byte) {
	if g.size == 0 {
		panic("gap's size is 0 - can't insert")
	}

	if g.cursor != g.offset {
		g.moveGap()
	}

	(*g.data)[g.offset] = char
	g.offset++
	g.cursor++
	g.size--
}

func (g *Gap) Delete() {
	if g.cursor != g.offset {
		g.moveGap()
	}

	if g.offset == 0 {
		return
	}

	g.offset--
	g.cursor--
	g.size++
}
