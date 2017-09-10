package demangler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMangledReader(t *testing.T) {
	sets := []struct {
		mangled string
	}{
		{"_imp__pthread_create"},
		{""},
	}
	for _, set := range sets {
		r := newMangledReader(set.mangled)
		assert.Equal(t, set.mangled, r.value)
		assert.Equal(t, len(set.mangled), r.Size())
		assert.Equal(t, len(set.mangled), r.Remaining())
		assert.Equal(t, 0, r.Offset())
	}
}

func TestMangledReaderSkip(t *testing.T) {
	r := newMangledReader("_imp__pthread_create")
	assert.Equal(t, 0, r.Offset())
	r.Skip(2)
	assert.Equal(t, 2, r.Offset())
	assert.Equal(t, "mp__pthread_create", r.Peek(r.Remaining()))
	r.Skip(-2)
	assert.Equal(t, "_imp__pthread_create", r.Peek(r.Remaining()))
	r.Skip(100)
	assert.Equal(t, 0, r.Offset())
	r.Skip(-1000)
	assert.Equal(t, 0, r.Offset())
}

func TestMangledReaderPeek(t *testing.T) {
	r := newMangledReader("_imp__pthread_create")
	assert.Equal(t, "_imp", r.Peek(4))
	assert.Equal(t, "", r.Peek(100))
	r.Skip(6)
	assert.Equal(t, "pthread", r.Peek(7))
}

func TestMangledReaderPeekRune(t *testing.T) {
	r := newMangledReader("_imp__pthread_create")
	assert.Equal(t, '_', r.PeekRune())
	r.Skip(6)
	assert.Equal(t, 'p', r.PeekRune())
}

func TestMangledReaderPeekRuneAt(t *testing.T) {
	r := newMangledReader("_imp__pthread_create")
	assert.Equal(t, 'i', r.PeekRuneAt(1))
	r.Skip(6)
	assert.Equal(t, 'r', r.PeekRuneAt(3))
}

func TestMangledReaderPeekEquals(t *testing.T) {
	r := newMangledReader("_imp__pthread_create")
	ok, len := r.PeekEquals("_imp", "__imp")
	assert.Equal(t, true, ok)
	assert.Equal(t, 4, len)
	ok, len = r.PeekEquals("_xmp", "__xmp")
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, len)
	r.Skip(6)
	ok, len = r.PeekEquals("pthread", "somethingorother")
	assert.Equal(t, true, ok)
	assert.Equal(t, 7, len)
}
