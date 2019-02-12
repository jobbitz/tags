package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testObject struct {
	ID   int
	Name string
	Role string `enum:"admin;moderator;user"`
}

func TestValidate(t *testing.T) {
	ass := assert.New(t)

	testObj := testObject{984, `yourmom`, `invalid`}
	ass.Error(Validate(&testObj))

	testObj.Role = `admin`
	ass.NoError(Validate(&testObj))
}
