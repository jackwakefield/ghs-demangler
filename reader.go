package demangler

type mangledReader struct {
	value string
	pos   int
}

func newMangledReader(value string) *mangledReader {
	return &mangledReader{
		value: value,
	}
}

// Size returns the length of the mangled name
func (m *mangledReader) Size() int {
	return len(m.value)
}

// Offset returns the current reading position of the mangled name
func (m *mangledReader) Offset() int {
	return m.pos
}

// Remaining returns the number of letters remaining in the mangled name
// after the offset
func (m *mangledReader) Remaining() int {
	return m.Size() - m.Offset()
}

// Skip increments the offset by the length specified
func (m *mangledReader) Skip(len int) {
	if m.Remaining() >= len {
		m.pos += len
		if m.pos < 0 {
			m.pos = 0
		}
	}
}

// Peek returns a string of the length specified from the current
// offset onwards without increasing the offset
// If the length specified is more than the remaining length
// an empty string is returned
func (m *mangledReader) Peek(len int) string {
	if m.Remaining() < len {
		return ""
	}
	return m.value[m.pos : m.pos+len]
}

// PeekEquals checks whether the text after the current offset matches
// any of the values given.
func (m *mangledReader) PeekEquals(values ...string) (bool, int) {
	for _, value := range values {
		if m.Peek(len(value)) == value {
			return true, len(value)
		}
	}
	return false, 0
}

// PeekRune returns a rune from the current offset without increasing
// the offset
// If the remaining length is smaller than 1, a rune with the
// value 0 is returned.
func (m *mangledReader) PeekRune() rune {
	return m.PeekRuneAt(0)
}

// PeekRuneAt returns a rune relative to the current offset
// If the remaining length is smaller than the relative offset, a rune
// with the value 0 is returned.
func (m *mangledReader) PeekRuneAt(relative int) rune {
	if m.Remaining() < relative+1 {
		return 0
	}
	return rune(m.value[m.pos+relative])
}
