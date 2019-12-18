// Copyright 2019 Job Stoit. All rights reserved.

package env

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// testConf is a test object
type testConf struct {
	Name string `json:"name" env:"NAME"`
	Sub  subObj `json:"sub"`
}

// subOnj
type subObj struct {
	Happy bool   `json:"happy" env:"HAPPY"`
	Data  []byte `json:"data"`
}

// TestMarshal tests the marshal function
func TestUnmarshal(t *testing.T) {
	as := assert.New(t)

	name, happy := `test`, true

	// Setting enviroment vars
	as.NoError(os.Setenv(`NAME`, name))
	as.NoError(os.Setenv(`HAPPY`, strconv.FormatBool(happy)))

	// Marshaling the test config
	conf := testConf{}
	as.NoError(Unmarshal(&conf))

	// Testing the correctness of the marshaler
	as.Equal(name, conf.Name)
	as.Equal(happy, conf.Sub.Happy)

	os.Unsetenv(`NAME`)
	os.Unsetenv(`HAPPY`)
}
