// Copyright 2020 Job Stoit. All rights reserved.

package flags

import (
	"os"
	"testing"
)

type testObj struct {
	Driver string   `flag:"driver"`
	Port   int      `flag:"p"`
	Demaon bool     `flag:"d"`
	Empty  string   `flag:"empty"`
	Args   []string `flag:"*"`
}

func TestParse(t *testing.T) {
	os.Args = []string{
		`app`,
		`-driver`, `postgres`,
		`-p`, `8080`,
		`-d`,
		`Arg1`, `Arg2`, `Arg3`, `Arg4`,
	}

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

	if len(obj.Args) != 4 {
		t.Errorf("Args not set: %q\n", obj.Args)
	}
}

//func ExampleParse() {
//	// Input: app -p 8080 -driver postgres arg1 arg2 arg3
//	os.Args = []string{`app`, `-p`, `8080`, `-driver`, `postgres`, `arg1`, `arg2`, `arg3`}
//
//	type Config struct {
//		Port               int      `flag:"p"`
//		DBConnectionString string   `flag:"cs"`
//		DBDriver           string   `flag:"driver"`
//		Args               []string `flag:"*"`
//	}
//
//	c := new(Config)
//	if err := Parse(c); err != nil {
//		panic(err)
//	}
//
//	fmt.Printf("port: %d, connection string: '%s', driver: '%s', args: %q\n", c.Port, c.DBConnectionString, c.DBDriver, c.Args)
//	// Output:
//	// port: 8080, connection string '', driver: 'postgres', args: ["arg1" "arg2" "arg3"]
//}
