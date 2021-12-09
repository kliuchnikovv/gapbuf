package gap

type Gap struct {
	offset int
	size   int

	Data   []byte
	cursor int
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

func (g *Gap) SetCursor(cursor int) {
	if cursor < 0 || cursor > g.sizeOfData() {
		return
	}

	g.cursor = cursor
}

func (g Gap) GetCursor() int {
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

	g.Data[g.offset] = char
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
	g.Data[g.offset] = 0
}
