// Copyright 2020 Job Stoit. All rights reserved.

package flags

import (
	"fmt"
	"os"
	"testing"

	"git.fuyu.moe/Fuyu/assert"
)

type testObj struct {
	Driver string   `flag:"driver"`
	Port   int      `flag:"p"`
	Demaon bool     `flag:"d"`
	Empty  string   `flag:"empty"`
	Args   []string `flag:"*"`
}

func TestParse(t *testing.T) {
	as := assert.New(t)
	os.Args = []string{
		`app`,
		`-driver`, `postgres`,
		`-p`, `8080`,
		`-d`,
		`Arg1`, `Arg2`, `Arg3`, `Arg4`,
	}

	obj := new(testObj)
	as.NoError(Parse(obj))

	as.Eq(`postgres`, obj.Driver)
	as.Eq(8080, obj.Port)
	as.Eq(true, obj.Demaon)
	as.Eq(``, obj.Empty)
	as.Eq(4, len(obj.Args))

	mainSet = Set{}
}

func ExampleParse() {
	// Input: app -p 8080 -driver postgres arg1 arg2 arg3
	os.Args = []string{`app`, `-p`, `8080`, `-driver`, `postgres`, `arg1`, `arg2`, `arg3`}

	type Config struct {
		Port               int      `flag:"p"`
		DBConnectionString string   `flag:"cs"`
		DBDriver           string   `flag:"driver"`
		Args               []string `flag:"*"`
	}

	c := new(Config)
	if err := Parse(c); err != nil {
		panic(err)
	}

	fmt.Printf("port: %d, connection string: '%s', driver: '%s', args: %q\n", c.Port, c.DBConnectionString, c.DBDriver, c.Args)
	// Output:
	// port: 8080, connection string: '', driver: 'postgres', args: ["arg1" "arg2" "arg3"]
}
