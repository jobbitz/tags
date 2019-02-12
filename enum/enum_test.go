package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testObject struct {
	ID     int
	Name   string
	Role   string  `enum:"admin;moderator;user"`
	Number float64 `enum:"1.2;2.0;4.6"`
}

func TestValidate(t *testing.T) {
	var err error
	ass := assert.New(t)

	testObj := testObject{984, `yourmom`, `invalid`, 0.0}
	err = Validate(&testObj)
	ass.Error(err)
	ass.Equal(Err{[]string{`admin`, `moderator`, `user`}, `invalid`}, err)

	testObj.Role = `admin`
	err = Validate(&testObj)
	ass.Error(err)
	ass.Equal(Err{[]string{`1.2`, `2.0`, `4.6`}, `0`}, err)

	testObj.Number = 4.6
	ass.NoError(Validate(&testObj))
}
