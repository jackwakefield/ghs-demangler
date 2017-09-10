package demangler

// Demangle takes a mangled Green Hill's C++ Compiler name and returns a
// demangled representation
// Mangled names with more than 255 characters will fail
func Demangle(mangled string) (string, error) {
	d := NewDemangler(mangled)
	return d.Demangle()
}

// Demangler provides an interface which takes a mangled name
// and returns a demangled representation
type Demangler interface {
	Demangle() (demangled string, err error)
}

type demangler struct {
	mangled string
	reader  *mangledReader

	dllImport   bool
	constructor int
	destructor  int
}

// NewDemangler creates an instance of the demangler for the given mangled name
func NewDemangler(mangled string) Demangler {
	return &demangler{
		mangled: mangled,
		reader:  newMangledReader(mangled),
	}
}

// Demangle attempts to demangle the mangled name
func (d *demangler) Demangle() (string, error) {
	if err := d.parsePrefix(); err != nil {
		return "", err
	}
	return "", nil
}

func (d *demangler) parsePrefix() error {
	d.dllImport = isDLLImport(d.reader)
	if !d.dllImport {
		if isGlobalConstructor(d.reader) {
			d.constructor = 2
		} else if isGlobalDestructor(d.reader) {
			d.destructor = 2
		}
	}
	return nil
}
