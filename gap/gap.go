package gap

type Gap struct {
	offset int
	size   int

	Data []byte
	// cursor int
}

func New(size int) *Gap {
	return &Gap{
		offset: 0,
		size:   size,
		Data:   make([]byte, size),
	}
}

func (g *Gap) Size() int {
	return g.size
}

func (g Gap) Offset() int {
	return g.offset
}

func (g *Gap) Insert(cursor int, char byte) {
	if g.size == 0 {
		panic("gap's size is 0 - can't insert")
	}

	if cursor != g.offset {
		g.moveGap(cursor)
	}

	g.Data[g.offset] = char
	g.offset++
	// g.cursor++
	g.size--
}

func (g *Gap) Delete(cursor int) {
	g.DeleteRange(cursor, 1)
}

func (g *Gap) DeleteRange(cursor, length int) {
	if g.sizeOfData() < length {
		return
	}

	if length == -1 {
		length = g.sizeOfData() - cursor
	}

	if length == 1 {
		g.moveGap(cursor)
	} else {
		g.moveGap(cursor + length)
	}

	if g.offset == 0 && length > 1 {
		return
	}

	for i := 0; i < length; i++ {
		g.offset--
		g.size++
		g.Data[g.offset] = 0
	}
}
