package demangler

var cplusMarkers = []rune{'.', '$'}

// findMarkerRune attempts to find a C++ marker at the current reader offset
func findMarkerRune(r *mangledReader) rune {
	return findMarkerRuneAt(r, 0)
}

// findMarkerRuneAt attempts to find a C++ marker relative to the
// current reader offset
func findMarkerRuneAt(r *mangledReader, offset int) rune {
	value := r.PeekRuneAt(offset)
	for _, marker := range cplusMarkers {
		if value == marker {
			return marker
		}
	}
	return 0
}
