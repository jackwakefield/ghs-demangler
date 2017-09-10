package demangler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDLLImport(t *testing.T) {
	sets := []struct {
		mangled string
		result  bool
		offset  int
	}{
		{"_imp__pthread_create", true, 6},
		{"__imp_pthread_create", true, 6},
		{"AllAwait__8NdSysFSSFv", false, 0},
		{"", false, 0},
	}
	for _, set := range sets {
		r := newMangledReader(set.mangled)
		assert.Equal(t, set.result, isDLLImport(r))
		assert.Equal(t, set.offset, r.Offset())
	}
}

func TestIsGlobalPrefix(t *testing.T) {
	sets := []struct {
		mangled string
		result  bool
		offset  int
	}{
		{"_GLOBAL_$I$tie__C3ios", true, 0},
		{"_GLOBAL_$D$tie__C3ios", true, 0},
		{"AllAwait__8NdSysFSSFv", false, 0},
		{"", false, 0},
	}
	for _, set := range sets {
		r := newMangledReader(set.mangled)
		assert.Equal(t, set.result, isGlobalPrefix(r))
		assert.Equal(t, set.offset, r.Offset())
	}
}

func TestIsGlobalConstructor(t *testing.T) {
	sets := []struct {
		mangled string
		result  bool
		offset  int
	}{
		{"_GLOBAL_$I$tie__C3ios", true, 11},
		{"_GLOBAL_*I*tie__C3ios", false, 0},
		{"_GLOBAL_$D$tie__C3ios", false, 0},
		{"AllAwait__8NdSysFSSFv", false, 0},
		{"", false, 0},
	}
	for _, set := range sets {
		r := newMangledReader(set.mangled)
		assert.Equal(t, set.result, isGlobalConstructor(r))
		assert.Equal(t, set.offset, r.Offset())
	}
}

func TestIsGlobalDestructor(t *testing.T) {
	sets := []struct {
		mangled string
		result  bool
		offset  int
	}{
		{"_GLOBAL_$I$tie__C3ios", false, 0},
		{"_GLOBAL_$D$tie__C3ios", true, 11},
		{"_GLOBAL_*D*tie__C3ios", false, 0},
		{"AllAwait__8NdSysFSSFv", false, 0},
		{"", false, 0},
	}
	for _, set := range sets {
		r := newMangledReader(set.mangled)
		assert.Equal(t, set.result, isGlobalDestructor(r))
		assert.Equal(t, set.offset, r.Offset())
	}
}
