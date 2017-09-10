package demangler

// isDLLImport determines whether the mangled name contains a symbol
// indicating it's being imported from a PE dynamic library
func isDLLImport(r *mangledReader) bool {
	ok, len := r.PeekEquals("_imp__", "__imp_")
	if ok {
		r.Skip(len)
	}
	return ok
}

func isGlobalPrefix(r *mangledReader) bool {
	ok, _ := r.PeekEquals("_GLOBAL_")
	return ok
}

func isGlobalPrefixWithRune(r *mangledReader, value rune) bool {
	if isGlobalPrefix(r) {
		marker := findMarkerRuneAt(r, 8)
		if marker == findMarkerRuneAt(r, 10) && r.PeekRuneAt(9) == value {
			return true
		}
	}
	return false
}

// isGlobalConstructor determines whether the mangled name is GNU global
// constructor to be executed at program init
func isGlobalConstructor(r *mangledReader) bool {
	if isGlobalPrefixWithRune(r, 'I') {
		r.Skip(11)
		return true
	}
	return false
}

// isGlobalConstructor determines whether the mangled name is GNU global
// destructor to be executed at program exit
func isGlobalDestructor(r *mangledReader) bool {
	if isGlobalPrefixWithRune(r, 'D') {
		r.Skip(11)
		return true
	}
	return false
}
