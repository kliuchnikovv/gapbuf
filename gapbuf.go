package gapbuf

import "bytes"

type GapBuffer struct {
	data []byte
	gap
}

func New(size int) *GapBuffer {
	var buffer = GapBuffer{
		data: make([]byte, size),
	}
	buffer.gap = *newGap(size, &buffer.data)

	return &buffer
}

func NewFromString(str string) *GapBuffer {
	var buffer = New(calculateNewSize(len(str)))
	buffer.Insert([]byte(str)...)
	return buffer
}

func (buffer *GapBuffer) Insert(bytes ...byte) {
	if int(buffer.cursor) != buffer.offset {
		buffer.MoveGap(int(buffer.cursor))
	}
	if buffer.size == 0 {
		// Extend buffer
		buffer.extend(len(bytes))
	} else if buffer.firstIndex()+len(bytes) >= buffer.lastIndex() {
		// Recursive call
		var extraBytes = bytes[buffer.size:]
		defer buffer.Insert(extraBytes...)

		bytes = bytes[:buffer.size]
	}

	for _, char := range bytes {
		buffer.setByte(char)
	}
}

func (buffer *GapBuffer) Delete(n int) {
	if buffer.firstIndex() == 0 || n <= 0 {
		return
	}
	if int(buffer.cursor) != buffer.offset {
		buffer.MoveGap(int(buffer.cursor))
	}
	buffer.offset--
	buffer.size++
}

func (buffer *GapBuffer) Split() []byte {
	var result = make([]byte, len(buffer.data)-int(buffer.cursor))
	copy(result, buffer.data[buffer.cursor:])
	buffer.data = buffer.data[:buffer.cursor]
	return result
}

func (buffer *GapBuffer) Bytes() []byte {
	return append(
		buffer.data[:buffer.firstIndex()],
		buffer.data[buffer.lastIndex():]...,
	)
}

func (buffer *GapBuffer) String() string {
	return string(buffer.Bytes())
}

func (buffer *GapBuffer) Size() int {
	return len(buffer.data) - buffer.size
}

func (buffer *GapBuffer) rawBytes() []byte {
	var result = make([]byte, len(buffer.data))
	copy(result, buffer.data)

	return append(
		result[:buffer.firstIndex()],
		append(
			bytes.Repeat([]byte{'.'}, buffer.size),
			result[buffer.lastIndex():]...,
		)...,
	)
}

func (buffer *GapBuffer) extend(extendSize int) {
	var (
		actualSize = len(buffer.data) + extendSize
		newSize    = calculateNewSize(actualSize)
	)

	buffer.size = newSize - actualSize

	var newSlice = make([]byte, actualSize)
	copy(newSlice, buffer.data[:buffer.offset])

	for i, char := range buffer.data[buffer.offset:] {
		newSlice[buffer.lastIndex()+i] = char
	}

	buffer.data = newSlice
	buffer.gap.data = &buffer.data
}
