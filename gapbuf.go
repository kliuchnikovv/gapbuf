package gapbuf

import (
	"github.com/kliuchnikovv/gapbuf/gap"
)

type GapBuffer struct {
	gap.Gap
}

func New(bytes ...byte) *GapBuffer {
	var (
		size   = calculateNewSize(len(bytes))
		buffer = GapBuffer{
			Gap: *gap.New(size),
		}
	)

	buffer.Insert(0, bytes...)

	return &buffer
}

func (buffer *GapBuffer) Insert(cursor int, bytes ...byte) {
	if cursor > len(buffer.String()) {
		cursor = len(buffer.String())
	}
	if buffer.Gap.Size() < len(bytes) {
		buffer.extend(len(bytes))
	}

	for _, char := range bytes {
		if char == 0 {
			break
		}
		buffer.Gap.Insert(cursor, char)
		cursor++
	}
}

func (buffer *GapBuffer) Split(cursor int) []byte {
	var (
		// cursor = buffer.GetCursor()
		data = buffer.Bytes()
	)

	buffer.Gap = *gap.New(len(buffer.Data))

	buffer.Insert(0, data[:cursor]...)

	// buffer.SetCursor(cursor)

	return data[cursor:]
}

func (buffer *GapBuffer) Bytes() []byte {

	if buffer.Gap.Size() == 0 {
		var result = make([]byte, len(buffer.Gap.Data))
		copy(result, buffer.Data)
		return result
	}

	if buffer.Offset() == buffer.Size() {
		var result = make([]byte, buffer.Offset())
		copy(result, buffer.Data[:buffer.Offset()])
		return result
	}

	var result = make([]byte, buffer.Size())
	copy(result, buffer.Data[:buffer.Gap.Offset()])

	for i, char := range buffer.Data[buffer.Gap.LastIndex():] {
		result[buffer.Gap.Offset()+i] = char
	}

	return result
}

func (buffer *GapBuffer) String() string {
	return string(buffer.Bytes())
}

func (buffer *GapBuffer) Size() int {
	var size = len(buffer.Data)
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
		actualSize = len(buffer.Data) + extendSize
		newSize    = calculateNewSize(actualSize)
		data       = buffer.Bytes()
	)

	buffer.Data = make([]byte, newSize)
	buffer.Gap = *gap.New(newSize)

	buffer.Insert(0, data...)
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
