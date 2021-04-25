// Copyright 2021 Job Stoit. All rights reserved.

package form

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testObj struct {
	Name string `form:"username"`
	Age  int    `form:"age"`
}

func TestEncode(t *testing.T) {
	as := assert.New(t)

	buf := new(bytes.Buffer)
	r, err := http.NewRequest(`POST`, `example.com`, buf)
	if err != nil {
		as.Fail(`failed to create the request`)
	}
	r.Form = url.Values{}

	name, age := `Bobby`, 25

	var sendObj testObj
	sendObj.Name = name
	sendObj.Age = age

	as.NoError(Encode(r, &sendObj))
	as.Equal(contentTypeVal, r.Header.Get(contentTypeKey))
	as.Equal(name, r.Form.Get(`username`))
	as.Equal(fmt.Sprint(age), r.Form.Get(`age`))
}

func TestDecode(t *testing.T) {
	as := assert.New(t)

	r, err := http.NewRequest(`POST`, `example.com`, nil)
	if err != nil {
		as.Fail(`failed to create the request`)
	}
	r.Form = url.Values{}

	name, age := `Blake`, 38
	r.Form.Add(`username`, name)
	r.Form.Add(`age`, fmt.Sprint(age))

	var recObj testObj
	as.Error(Decode(r, &recObj))
	as.Equal(``, recObj.Name)
	as.Equal(0, recObj.Age)

	r.Header.Set(contentTypeKey, contentTypeVal)
	as.NoError(Decode(r, &recObj))
	as.Equal(name, recObj.Name)
	as.Equal(age, recObj.Age)
}
