package gapbuf

import (
	"github.com/KlyuchnikovV/gapbuf/gap"
)

type GapBuffer struct {
	data []byte
	gap.Gap
}

func New(bytes ...byte) *GapBuffer {
	var (
		size   = calculateNewSize(len(bytes))
		buffer = GapBuffer{
			data: make([]byte, size),
		}
	)
	buffer.Gap = *gap.New(size, &buffer.data)

	buffer.Insert(bytes...)

	return &buffer
}

func (buffer *GapBuffer) Insert(bytes ...byte) {
	if buffer.Gap.Size() < len(bytes) {
		buffer.extend(len(bytes))
	}

	for _, char := range bytes {
		if char == 0 {
			break
		}
		buffer.Gap.Insert(char)
	}
}

func (buffer *GapBuffer) Split() []byte {
	var (
		cursor = buffer.GetCursor()
		data   = buffer.Bytes()
	)

	buffer.Gap = *gap.New(len(buffer.data), &buffer.data)

	buffer.Insert(data[:cursor]...)

	buffer.SetCursor(cursor)

	return data[cursor:]
}

func (buffer *GapBuffer) Bytes() []byte {
	if buffer.Gap.Size() == 0 {
		return buffer.data
	}

	if buffer.Offset() == buffer.Size() {
		return buffer.data[:buffer.Offset()]
	}

	return append(
		buffer.data[:buffer.Gap.Offset()],
		buffer.data[buffer.Gap.LastIndex():]...,
	)
}

func (buffer *GapBuffer) String() string {
	return string(buffer.Bytes())
}

func (buffer *GapBuffer) Size() int {
	var size = len(buffer.data)
	if size-buffer.Gap.Size() < 0 {
		return size
	}
	return size - buffer.Gap.Size()
}

func (buffer *GapBuffer) extend(extendSize int) {
	if extendSize == 0 {
		return
	}
	var (
		cursor     = buffer.GetCursor()
		actualSize = len(buffer.data) + extendSize
		newSize    = calculateNewSize(actualSize)
		data       = buffer.Bytes()
	)

	buffer.data = make([]byte, newSize)
	buffer.Gap = *gap.New(newSize, &buffer.data)

	buffer.Insert(data...)
	buffer.SetCursor(cursor)
}

func calculateNewSize(size int) int {
	var result = size
	switch {
	case size == 0:
		result = 0
	case size <= 10:
		result = 10
	case size <= 20:
		result = 20
	case size <= 40:
		result = 40
	default:
		result += 40
	}
	return result
}
