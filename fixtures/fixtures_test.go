package fixtures

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/decaf-emu/ghs-demangle"
	"github.com/stretchr/testify/assert"
)

func TestFixtures(t *testing.T) {
	items := make(map[string]string)
	data, err := ioutil.ReadFile("fixtures/mp10-tenst.json")
	if err != nil {
		t.Error(err)
	}
	if err := json.Unmarshal(data, &items); err != nil {
		t.Error(err)
	}
	successful := 0
	for mangled, expected := range items {
		demangled, err := demangler.Demangle(mangled)
		if assert.NoError(t, err) && assert.Equal(t, expected, demangled) {
			successful++
		}
	}
	t.Logf("%d/%d passes", successful, len(items))
}
