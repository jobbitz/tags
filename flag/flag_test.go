package flag

import (
	f "flag"
	"testing"
)

type testStr struct {
	Verbose bool   `flag:"v"`
	Name    string `flag:"name"`
}

func TestUnmarshal(t *testing.T) {
	obj := testStr{}
	f.Set(`v`, ``)
	f.Set(`name`, `hello`)

	Unmarshal(&obj)

	if obj.Name != `hello` {
		t.Errorf("name not parsed: %s\n", obj.Name)
	}

	if obj.Verbose == false {
		t.Errorf("verbose not active : %v\n", obj.Verbose)
	}
}
