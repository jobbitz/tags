// Copyright 2021 Job Stoit. All rights reserved.

package form

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"git.fuyu.moe/Fuyu/assert"
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
		t.Error(`failed to create the request`)
		t.Fail()
	}
	r.Form = url.Values{}

	name, age := `Bobby`, 25

	var sendObj testObj
	sendObj.Name = name
	sendObj.Age = age

	as.NoError(Marshal(r, &sendObj))
	as.Eq(contentTypeVal, r.Header.Get(contentTypeKey))
	as.Eq(name, r.Form.Get(`username`))
	as.Eq(fmt.Sprint(age), r.Form.Get(`age`))
}

func TestDecode(t *testing.T) {
	as := assert.New(t)

	r, err := http.NewRequest(`POST`, `example.com`, nil)
	if err != nil {
		t.Error(`failed to create the request`)
		t.Fail()
	}
	r.Form = url.Values{}

	name, age := `Blake`, 38
	r.Form.Add(`username`, name)
	r.Form.Add(`age`, fmt.Sprint(age))

	var recObj testObj
	r.Header.Set(contentTypeKey, contentTypeVal)
	as.NoError(Unmarshal(r, &recObj))
	as.Eq(name, recObj.Name)
	as.Eq(age, recObj.Age)
}
