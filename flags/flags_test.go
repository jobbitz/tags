// Copyright 2020 Job Stoit. All rights reserved.

package flags

import (
	"os"
	"testing"
)

type testObj struct {
	Driver string `flag:"driver"`
	Port   int    `flag:"p"`
	Demaon bool   `flag:"d"`
	Empty  string `flag:"empty"`
}

func TestUnmarshal(t *testing.T) {
	os.Args = []string{`app`, `-driver`, `postgres`, `-p`, `8080`, `-d`}

	obj := new(testObj)
	if err := Parse(obj); err != nil {
		t.Error(err)
	}

	if obj.Driver != `postgres` {
		t.Error(`driver not parsed`)
	}

	if obj.Port != 8080 {
		t.Error(`port not parsed`)
	}

	if obj.Demaon != true {
		t.Error(`deamon not parsed`)
	}

	if obj.Empty != `` {
		t.Error(`empty should not be set`)
	}

}
