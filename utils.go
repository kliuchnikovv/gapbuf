package gapbuf

type cursor int

func (c *cursor) SetCursor(cur int) {
	*c = cursor(cur)
}

type gap struct {
	offset int
	size   int

	data *[]byte
}

func newGap(size int, pointer *[]byte) *gap {
	return &gap{
		offset: 0,
		size:   size,
		data:   pointer,
	}
}

func (g gap) firstIndex() int {
	return g.offset
}

func (g gap) lastIndex() int {
	return g.offset + g.size
}

func (g *gap) MoveGap(cursor int) {
	if cursor < g.offset {
		for i, char := range (*g.data)[cursor:g.offset] {
			(*g.data)[cursor+g.size+i] = char
		}
	} else {
		for i, char := range (*g.data)[g.lastIndex() : g.lastIndex()+g.size] {
			(*g.data)[g.offset+i] = char
		}
	}
	g.offset = cursor
}

func (g *gap) setByte(char byte) {
	(*g.data)[g.offset] = char
	g.offset++
	g.size--
}

func calculateNewSize(size int) int {
	var result int
	switch {
	case size == 0:
		result = 10
	case size <= 10:
		result = 20
	case size <= 20:
		result = 40
	default:
		result += 40
	}
	return result
}
