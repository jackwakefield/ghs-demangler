package demangler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMarkerRune(t *testing.T) {
	r := newMangledReader(".")
	result := findMarkerRune(r)
	assert.Equal(t, '.', result)

	r = newMangledReader("$")
	result = findMarkerRune(r)
	assert.Equal(t, '$', result)

	r = newMangledReader("test")
	result = findMarkerRune(r)
	assert.Equal(t, rune(0), result)
}

func TestFindMarkerRuneAt(t *testing.T) {
	r := newMangledReader("abc.def$")
	result := findMarkerRuneAt(r, 3)
	assert.Equal(t, '.', result)

	result = findMarkerRuneAt(r, 7)
	assert.Equal(t, '$', result)

	result = findMarkerRuneAt(r, 2)
	assert.Equal(t, rune(0), result)

	result = findMarkerRuneAt(r, 100)
	assert.Equal(t, rune(0), result)
}
